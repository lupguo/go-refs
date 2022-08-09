package synclib

import (
	"sync"
	"testing"
)

var l sync.Mutex
var a string

func f() {
	a = "hello, world"
	l.Unlock() // release lock
}

func TestLockMuchTims(t *testing.T) {
	// l.Lock() // no block, lock
	go f()
	l.Lock() // wait unlock
	// defer l.Unlock()
	println(a)
}
