package goroutine

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestWait(t *testing.T) {
	wg := &sync.WaitGroup{}
	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func() {
			// defer func() {
			// 	wg.Done()
			// }()
			defer wg.Done()
			time.Sleep(2 * time.Second)
		}()
	}
	wg.Wait()
	fmt.Print("Done")
}
