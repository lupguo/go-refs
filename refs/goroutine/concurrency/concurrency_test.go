package concurrency

import (
	"fmt"
	"runtime"
	"sync"
	"testing"
	"time"
)

func TestConcurrency(t *testing.T) {
	// create new channel of type string
	ch := make(chan string)

	// start new anonymous goroutine
	go func() {
		time.Sleep(time.Second)
		// send "Hello World" to channel
		ch <- "Hello World"
	}()
	// read from channel
	msg, ok := <-ch
	fmt.Printf("msg='%s', ok='%v'\n", msg, ok)
}

func TestCreateGoroutine(t *testing.T) {
	mult := func(x, y int) {
		fmt.Printf("%d * %d = %d\n", x, y, x*y)
	}
	for i := 0; i < 10; i++ {
		go mult(1, 2) // first execution, non-blocking
	}
}

func TestWatiGroup(t *testing.T) {
	var wg sync.WaitGroup // 1

	routine := func(i int) {
		defer wg.Done() // 3
		fmt.Printf("routine %v finished\n", i)
	}

	wg.Add(10) // 2
	for i := 0; i < 10; i++ {
		go routine(i) // *
	}
	wg.Wait() // 4
	fmt.Println("main finished")
}

func TestLimitConcurrency(t *testing.T) {
	var (
		semaphoreSize = runtime.NumCPU()

		mu                 sync.Mutex
		totalTasks         int
		curConcurrentTasks int
		maxConcurrentTasks int
	)

	var timeConsumingTask = func() {
		mu.Lock()
		totalTasks++
		curConcurrentTasks++
		if curConcurrentTasks > maxConcurrentTasks {
			maxConcurrentTasks = curConcurrentTasks
		}
		mu.Unlock()

		// in real system this would be a CPU intensive operation
		time.Sleep(10 * time.Millisecond)

		mu.Lock()
		curConcurrentTasks--
		mu.Unlock()
	}

	// 并发度控制
	var sem = make(chan struct{}, semaphoreSize)
	var wg sync.WaitGroup
	for i := 0; i < 32; i++ {
		// acquire semaphore
		sem <- struct{}{}
		wg.Add(1)

		go func() {
			timeConsumingTask()
			// release semaphore
			<-sem
			wg.Done()
		}()
	}
	// wait for all task to finish
	wg.Wait()
	fmt.Printf("total tasks         : %d\n", totalTasks)
	fmt.Printf("max concurrent tasks: %d\n", maxConcurrentTasks)
}
