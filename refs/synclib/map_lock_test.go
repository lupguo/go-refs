package synclib

import (
	"fmt"
	"sync"
	"testing"
)

func Lock(key string) {
	locks := make(map[string]*sync.RWMutex)

	fn := func() {
		kmu, ok := locks[key]
		if !ok { // 锁不存在，初始化一个设置
			kmu = &sync.RWMutex{}
			locks[key] = kmu
		}
		kmu.Lock()
	}

	fn()
	kmu, ok := locks[key]
	if ok {
		fmt.Println("unlock")
		kmu.Unlock()
	} else {
		fmt.Println("unlock no locks")
	}
}

func TestLock(t *testing.T) {
	Lock("hello")

	Lock("hello")
}
