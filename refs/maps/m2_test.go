package maps

import (
	"testing"
	"time"
)

var m = map[int]int{1: 100}

func TestConrrRead(t *testing.T) {

	go read()
	time.Sleep(time.Second)

	// go write()
}

func write() {
	for {
		m[1] = 1
	}
}

func read() {
	for {
		_ = m[1]
	}
}
