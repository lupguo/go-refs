package channel

import (
	"testing"
	"time"
)

func TestPrintChan(t *testing.T) {
	ch := make(chan int)
	go func(c chan int) {
		t.Log(<-c)
	}(ch)
	ch<-5
	time.Sleep(time.Second)
}
