package syncxmap

import (
	"fmt"
	"sync"
	"testing"
)

func TestNormalMap(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 1e3; i++ {
		wg.Add(1)
		i := i // go routine 循环陷阱
		go func() {
			mapAddObject(fmt.Sprintf("num:%v", i), i)
			wg.Done()
		}()
	}
	wg.Wait()

	for key, val := range normalStorage {
		var num int
		n, err := fmt.Sscanf(key, "num:%d", &num)
		if err != nil {
			t.Error(err)
		}
		t.Logf("n=%v, num=%d", n, num)

		if val != num {
			t.Errorf("got num=%v, but want %v", num, val)
		}
	}

	for i := 0; i < 1e3; i++ {
		key := fmt.Sprintf("num:%d", i)
		if got := normalStorage[key]; got != i {
			t.Errorf("normal stroage key got err, got %v, but want %v", got, i)
		}
	}
}

func TestBenchSyncMap(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 1e2; i++ {
		wg.Add(1)
		i := i // go routine 循环陷阱
		go func() {
			syncMapAddObject(fmt.Sprintf("num:%v", i), i)
			wg.Done()
		}()
	}
	wg.Wait()

	// 校验
	syncStorage.Range(func(key, value any) bool {
		// skey
		skey, ok := key.(string)
		if !ok {
			t.Errorf("skey got err")
		}
		var num int
		n, err := fmt.Sscanf(skey, "num:%d", &num)
		if err != nil {
			t.Error(err)
			return false
		}
		t.Logf("n=%v, num=%d", n, num)

		// sval
		sval, ok := value.(int)
		if !ok {
			t.Errorf("sval got err")
			return false
		}
		if num != sval {
			t.Errorf("got num=%v, but want %v", num, sval)
		}

		return true
	})

}

var syncStorage sync.Map

// sync map 存取
func syncMapAddObject(key string, val int) {
	syncStorage.Store(key, val)
}

var normalStorage map[string]int
var locker sync.RWMutex
var once sync.Once

func mapAddObject(key string, val int) {
	// if normalStorage == nil { // 非并发安全
	// 	fmt.Println("init normal map")
	// 	normalStorage = make(map[string]int)
	// }
	once.Do(func() {
		fmt.Println("init normal map")
		normalStorage = make(map[string]int)
	})

	// 增加读写锁
	locker.Lock()
	defer locker.Unlock()
	normalStorage[key] = val
}
