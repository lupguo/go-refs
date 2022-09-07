package syncx

import (
	"sync"
	"testing"
	"time"
)

// once
var once sync.Once

// 初始化默认值
var defaultVal int

func CreateInstance(id int, t *testing.T) {
	if defaultVal != 0 {
		t.Logf("#%d found default val has init, return defaulVal=%d", id, defaultVal)
		return
	}

	t.Logf("#%d create begin", id)
	once.Do(func() { // 多个协程并发执行，仅有一个会做初始化，同时内部会基于sync.Once内部互斥锁机制排队等待，一旦初始化成功则直接返回
		t.Logf("#%d start init default val...", id)
		time.Sleep(500 * time.Millisecond)
		defaultVal = 9
	})

	t.Logf("#%d create done, got defaultVal=%d", id, defaultVal)
	if defaultVal == 0 {
		t.Errorf("oops, #%d got zere defaultVal=%d", id, defaultVal)
	}
}

func TestName(t *testing.T) {
	for i := 0; i < 100; i++ {
		go func(id int) {
			CreateInstance(id, t)
		}(i)
	}

	time.Sleep(1 * time.Second)
}

func TestNilOnce(t *testing.T) {
	var o *sync.Once // panic, o的值为nil
	o.Do(func() {
		t.Logf("hello")
	})
}

func TestNilOnce2(t *testing.T) {
	var o sync.Once // ok, o的值为初始化的
	o.Do(func() {
		t.Logf("hello")
	})
}
