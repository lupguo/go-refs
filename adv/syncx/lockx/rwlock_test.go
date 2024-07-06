package lockx

import (
	"sync"
	"testing"
	"time"
)

func TestCase1(t *testing.T) {
	rwl := &sync.RWMutex{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 写锁协程
	g1 := func() {
		// 写锁
		rwl.Lock()
		defer rwl.Unlock()
		t.Log("get write lock succ")
		time.Sleep(1 * time.Second)
		wg.Done()
	}
	go g1()

	// 加读锁
	for i := 0; i < 5; i++ {
		wg.Add(1)
		n := i
		go func() {
			defer wg.Done()
			rwl.RLock()
			defer rwl.RUnlock()
			t.Logf("#[%d] running, &i=%p, &n=%p", n, &i, &n)
		}()
	}

	// 等待读、写锁协程同步完成
	wg.Wait()
}

func TestCase2(t *testing.T) {
	wl := &sync.Mutex{}
	wg := sync.WaitGroup{}
	wg.Add(1)
	// 写锁协程
	g1 := func() {
		// 写锁
		wl.Lock()
		defer wl.Unlock()
		t.Log("get mutex lock succ")
		time.Sleep(1 * time.Second)
		wg.Done()
	}
	go g1()

	// 加读锁
	for i := 0; i < 5; i++ {
		wg.Add(1)
		n := i
		go func() {
			defer wg.Done()
			wl.Lock()
			defer wl.Unlock()
			t.Logf("#[%d] running, get lock succ, &i=%p, &n=%p", n, &i, &n)
		}()
	}

	// 等待读、写锁协程同步完成
	wg.Wait()
}
