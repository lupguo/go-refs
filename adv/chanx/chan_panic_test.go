package chanx

import (
	"testing"
)

// panic: close of closed channel
func TestPanic1(t *testing.T) {
	ch := make(chan int)
	close(ch)
	close(ch)
}

// panic: send on closed channel
func TestPanic2(t *testing.T) {
	ch := make(chan int)
	go func() {
		<-ch
	}()
	close(ch)
	ch <- 1 // panic: send on closed channel
}

// panic 无接收协程就开始发送
// fatal error: all goroutines are asleep - deadlock!
func TestPanic3(t *testing.T) {
	ch := make(chan int)
	ch <- 1
}

// panic 无发送协程就开始接收
// fatal error: all goroutines are asleep - deadlock!
func TestPanic4(t *testing.T) {
	ch := make(chan int, 1)
	<-ch
}

// panic : close of nil channel
func TestPanic5(t *testing.T) {
	var ch chan int
	close(ch)
}

// panic: chan receive (nil chan)
func TestPanic6(t *testing.T) {
	var ch chan int
	go func() {
		ch <- 1
	}()
	<-ch
}
