package memmod

import (
	"testing"
	"time"
)

func TestChanUnBufferPanic(t *testing.T) {
	ch := make(chan string, 10)
	go func() {
		time.Sleep(1 * time.Second)
		<-ch
	}()
	for i := 0; i <= 10; i++ {
		ch <- "hey"
	}
}

func TestChanUnBuffer(t *testing.T) {
	ch := make(chan string)
	go func() {
		ch <- "hey"
	}()
	if v, ok := <-ch; ok {
		t.Logf("v=>%v", v)
	}
}
