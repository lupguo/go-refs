package lockx

import (
	"fmt"
	"sync"
	"sync/atomic"
	"testing"
)

// 乐观锁
type OptimisticLock struct {
	value int64
}

// 乐观锁：通过for循环spin自旋，尝试atomic原子操作+cas实现乐观更新
func (ol *OptimisticLock) Increment() {
	for {
		current := atomic.LoadInt64(&ol.value)
		next := current + 1
		if atomic.CompareAndSwapInt64(&ol.value, current, next) {
			break
		}
	}
}

func (ol *OptimisticLock) GetValue() int64 {
	return atomic.LoadInt64(&ol.value)
}

// 悲观锁
type PessimisticLock struct {
	value int
	mu    sync.Mutex
}

// 悲观锁：直接加锁，再处理竟态资源
func (pl *PessimisticLock) Increment() {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	pl.value++
}

// 悲观锁：直接加锁，再处理竟态资源
func (pl *PessimisticLock) GetValue() int {
	pl.mu.Lock()
	defer pl.mu.Unlock()
	return pl.value
}

// 乐观锁测试
func TestOptimisticLock(t *testing.T) {
	var wg sync.WaitGroup
	ol := OptimisticLock{} // 竟态资源

	// 100个并发协程
	for i := 0; i < 100; i++ {
		wg.Add(1)

		// 并发加数
		go func() {
			ol.Increment() // 竟态资源更新冲突（乐观处理）
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", ol.GetValue())
}

// 悲观锁测试
func TestPessimisticLock(t *testing.T) {
	var wg sync.WaitGroup
	pl := PessimisticLock{} // 竟态资源

	for i := 0; i < 100; i++ {
		wg.Add(1)
		go func() {
			pl.Increment() // 竟态资更新处理（悲观处理）
			wg.Done()
		}()
	}

	wg.Wait()
	fmt.Println("Final value:", pl.GetValue())
}
