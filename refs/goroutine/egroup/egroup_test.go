package egroup

import (
	"sync"
	"testing"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

// 4950
func TestSum(t *testing.T) {
	sum := 0
	for i := 0; i < 100; i++ {
		sum += i
	}
	t.Log(sum)
}

// 无锁版，存在sum竟态问题、闭包问题 (5577)
func TestErrGroupV1(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	for i := 0; i < 100; i++ {
		egp.Go(func() error {
			sum += i
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

// 有锁版，sum加锁，但存在闭包问题 (8382)
func TestErrGroupV2(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	lock := sync.Mutex{}
	for i := 0; i < 100; i++ {
		egp.Go(func() error {
			lock.Lock()
			defer lock.Unlock()
			sum += i
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

// sum加锁，赋值解决闭包问题(4950)
func TestErrGroupV3(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	lock := sync.Mutex{}
	for i := 0; i < 100; i++ {
		j := i
		egp.Go(func() error {
			lock.Lock()
			defer lock.Unlock()
			sum += j
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

// sum加锁，赋值解决闭包问题, i=i尝试，也是ok的（）
func TestErrGroupV4(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	lock := sync.Mutex{}
	for i := 0; i < 100; i++ {
		i := i // 相当于重新申请了临时变量i，没有了闭包问题
		egp.Go(func() error {
			t.Logf("i=%d", i)
			lock.Lock()
			defer lock.Unlock()
			sum += i
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

func sumTotal(sum *int, n int) {
	*sum += n
}

// 函数版本，有问题(&sum、i变量均存在竟态问题）
func TestErrGroupFnV1(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	for i := 0; i < 100; i++ {
		egp.Go(func() error {
			sumTotal(&sum, i)
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

// 函数版本，有问题(&sum存在竟态问题）
func TestErrGroupFnV2(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	for i := 0; i < 100; i++ {
		i := i
		egp.Go(func() error {
			sumTotal(&sum, i)
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

// 函数加锁版本(5604)，i还是存在闭包(i存在竟态问题，&sum通过锁解决了并发问题）
func TestErrGroupFnV3(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	lock := sync.Mutex{}
	for i := 0; i < 100; i++ {
		egp.Go(func() error {
			lock.Lock()
			t.Logf("i=%d", i) // i 是值拷贝
			sumTotal(&sum, i)
			lock.Unlock()
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

// 函数加锁版本，重新赋值解决闭包(4950) - OK
func TestErrGroupFnV4(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	lock := sync.Mutex{}
	for i := 0; i < 100; i++ {
		j := i
		egp.Go(func() error {
			lock.Lock()
			t.Logf("j=%d", j)
			sumTotal(&sum, j)
			lock.Unlock()
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

// 函数加锁版本，重新赋值解决闭包, i:=i
func TestErrGroupFnV5(t *testing.T) {
	egp := errgroup.Group{}
	sum := 0
	lock := sync.Mutex{}
	for i := 0; i < 100; i++ {
		i := i
		egp.Go(func() error {
			lock.Lock()
			t.Logf("i=%d", i)
			sumTotal(&sum, i)
			lock.Unlock()
			return nil
		})
	}
	if err := egp.Wait(); err != nil {
		t.Error(err)
	}
	t.Log(sum)
}

func TestEgroupExecOver(t *testing.T) {
	eg := errgroup.Group{}
	eg.Go(func() error {
		time.Sleep(1000 * time.Millisecond)
		t.Log("business ok")
		return nil
	})
	eg.Go(func() error {
		return errors.New("hack error")
	})

	// eg.Wait会等待所有eg.Go()方法返回后才返回
	if err := eg.Wait(); err != nil {
		t.Errorf("got err: %s", err)
	}
}
