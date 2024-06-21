package bounded

import (
	"fmt"
	"sort"
	"testing"
	"time"
)

// 160~241ms
func TestMD5All(t *testing.T) {
	// 程序执行时间
	start := time.Now()
	// 注意: 这里需要有个匿名函数包裹，不能直接用 defer t.Log(time.Since(start))
	defer func() {
		t.Log(time.Since(start))
	}()

	// Calculate the MD5 sum of all files under the specified directory,
	// then print the results sorted by path name.

	root := `/data/python`
	m, err := MD5All(root)
	if err != nil {
		fmt.Println(err)
		return
	}
	var paths []string
	for path := range m {
		paths = append(paths, path)
	}
	sort.Strings(paths)
}
