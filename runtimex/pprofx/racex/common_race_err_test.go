package racex

import (
	"testing"
)

func TestParallelWrite(t *testing.T) {
	// ParallelWrite([]byte("hello"))
}

func TestChanRace(t *testing.T) {
	c := make(chan struct{}) // or buffered channel

	go func() { c <- struct{}{} }()
	<-c
	close(c)
}
