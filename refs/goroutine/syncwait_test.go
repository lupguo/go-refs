package goroutine

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
)

func TestWait(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			t.Logf("goroutine #%d", i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
	fmt.Print("Done")
}

func TestWaitSema(t *testing.T) {
	over := make(chan struct{})
	// 轮询检测goroutine数量
	asyncShowGoroutineNum(t, over, 100*time.Millisecond)
	defer close(over)

	// 通过sema信号量控制并发，通过wg进行同步
	sema := make(chan struct{}, 10)
	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		sema <- struct{}{}
		go func(i int) {
			defer func() {
				<-sema
				wg.Done()
			}()
			t.Logf("goroutine #%d", i)
			time.Sleep(time.Second)
		}(i)
	}
	wg.Wait()
	t.Log("Done")
}

func asyncShowGoroutineNum(t *testing.T, over chan struct{}, interval time.Duration) {
	go func() {
		for {
			select {
			case <-time.Tick(interval):
				t.Logf("goroutine cnt(%d)", runtime.NumGoroutine())
			case <-over:
				t.Logf("exist test")
				return
			}
		}
	}()
}

func TestErrorGroup(t *testing.T) {
	// 轮询检测goroutine数量
	over := make(chan struct{})
	asyncShowGoroutineNum(t, over, 100*time.Millisecond)
	defer close(over)

	// errgroup 控制并发
	egp := errgroup.Group{}
	egp.SetLimit(2)
	for i := 0; i < 100; i++ {
		i = i
		egp.Go(func() error {
			t.Logf("goroutine #%d", i)
			time.Sleep(time.Second)
			return nil
		})
	}

	if err := egp.Wait(); err != nil {
		t.Errorf("egp got err: %s", err)
	}

	t.Log("Done")
}
