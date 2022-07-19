package goroutine

import (
	"fmt"
	"testing"
	"time"
)

func TestMutilGo(t *testing.T) {
	go func() {
		go func() {
			defer func() {
				if r := recover(); r != nil {
					t.Logf("recover: %+v", r)
				}
			}()
			time.Sleep(2 * time.Second)
			panic("fun1 panic")
		}()

		t.Logf("go fun1 exec done")
	}()

	go func() {
		sum := 0

		go func() {
			for i := 0; i < 100; i++ {
				time.Sleep(50 * time.Millisecond)
				sum += i
			}

			t.Logf("go fun3 exec done")
		}()

		go func() {
			for {
				select {
				case <-time.Tick(100 * time.Millisecond):
					t.Logf("sum=>%d", sum)
				}
			}
		}()

		t.Logf("go fun2 exec done")
	}()

	time.Sleep(10 * time.Second)

	t.Logf("main exit")
}

func TestG(t *testing.T) {
	// 合起来写
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}
