package chanx

import (
	"testing"
	"time"
)

func TestSelect(t *testing.T) {
	ch := make(chan int)
	go func() {
		i := 0
		for {
			select {
			case <-time.After(time.Second):
				ch <- i
			}
			i++
		}
	}()
	go func() {
		for v := range ch {
			t.Logf("got v %v", v)
		}
	}()

	time.Sleep(3 * time.Second)
}
