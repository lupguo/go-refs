// Package bounded 整个bounded流程
// 1. filepath.Walk()遍历目录，将目录树路径结果通过chan返回
// 2. 创建一个签名函数 digester()，通过从chan接收path数据，读取文件内容，进行摘要签名
// 3.
package bounded

import (
	"crypto/md5"
	"os"
	"path/filepath"
	"sync"

	"github.com/pkg/errors"
)

type result struct {
	path string
	sum  [md5.Size]byte
	err  error
}

// 启动一个非阻塞协程，walk遍历整个root路径下的内容
// 通过 info.Mode().IsRegular()识别是否标准文件
// 通过 select检测是否整个walkFiles需要中断退出，避免协程资源泄露
func walkFiles(done <-chan struct{}, root string) (<-chan string, <-chan error) {
	paths := make(chan string)
	errc := make(chan error, 1)
	go func() {
		// Close the paths channel after Walk returns.
		defer close(paths)
		// No select needed for this send, since errc is buffered.
		errc <- filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}
			if !info.Mode().IsRegular() {
				return nil
			}
			select {
			case paths <- path:
			case <-done:
				return errors.New("walk canceled")
			}
			return nil
		})
	}()
	return paths, errc
}

func digester(done <-chan struct{}, paths <-chan string, c chan<- result) {
	for path := range paths {
		data, err := os.ReadFile(path)
		select {
		case c <- result{path, md5.Sum(data), err}:
		case <-done:
			return
		}
	}
}

func MD5All(root string) (map[string][md5.Size]byte, error) {
	// MD5All closes the done channel when it returns; it may do so before
	// receiving all the values from c and errc.
	done := make(chan struct{})
	defer close(done)

	// 非阻塞的并发遍历子文件，进行子文件IO读取，计算sum后返回
	paths, errc := walkFiles(done, root)

	// 并发签名paths内容，将结果回写到c中
	c := make(chan result)
	var wg sync.WaitGroup
	const numDigesters = 20
	wg.Add(numDigesters)
	for i := 0; i < numDigesters; i++ {
		go func() {
			digester(done, paths, c)
			wg.Done()
		}()
	}

	go func() {
		// 当所有前面完结后，关闭c
		wg.Wait()
		close(c)
	}()

	// 最后，从c中读取摘要前面
	m := make(map[string][md5.Size]byte)
	for r := range c {
		if r.err != nil {
			return nil, r.err
		}
		m[r.path] = r.sum
	}

	// Check whether the Walk failed.
	if err := <-errc; err != nil {
		return nil, err
	}
	return m, nil

}
