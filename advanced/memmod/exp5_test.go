package memmod

import (
	"sync"
	"testing"
)

func TestExp5Lock(t *testing.T) {
	l := sync.Mutex{}
	l.Lock()
	l.Lock() // panic
}

func TestUnlock(t *testing.T) {
	var l sync.Mutex
	l.Unlock()
}

// 输出a=0
//	无法保证go协程内的a=1赋值在log之前
func TestLockExp1(t *testing.T) {
	a := 0
	f := func() {
		a = 1
	}
	go f()
	t.Logf("a=>%v", a)
}

// 输出a=1
//	通过加锁确保顺序
func TestLockExp2(t *testing.T) {
	m := sync.Mutex{}
	a := 0
	f := func() {
		a = 1
		m.Unlock() // unlock1
	}
	m.Lock() // lock1
	go f()
	m.Lock()   // lock2 block, wait lock1 release by unlock1
	m.Unlock() // unlock2

	t.Logf("a=>%v", a)
}
