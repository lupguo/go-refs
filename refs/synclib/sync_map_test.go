package synclib

import (
	"sync"
	"testing"
)

func TestSyncMap(t *testing.T) {
	syncMap := &sync.Map{}

	// 1. 设置值
	syncMap.Store("user1", "Terry")
	syncMap.Store("user2", "Gold")

	// 2. 循环值
	syncMap.Range(func(key, value interface{}) bool {
		// range出来还是interface
		t.Logf("key=>%v, val=>%v", key, value)
		t.Logf("print again, key=>%s, val=>%s", key, value)
		return true
	})

	// 3. 获取单个值
	if v, ok := syncMap.Load("user1"); ok {
		t.Logf("user1 exist =>%v", v)
	}
	if _, ok := syncMap.Load("user3"); !ok {
		t.Logf("user3 not exist")
	}

	// 4. 加载并删除值
	if v, ok := syncMap.LoadAndDelete("user1"); ok {
		t.Logf("user1 delete success, got user1 %v", v)

		// load again
		if _, ok := syncMap.LoadAndDelete("user1"); !ok {
			t.Logf("user1 now got empty...")
		}
	}

	// 5. 加载并存储
	if _, ok:= syncMap.LoadOrStore("user1", "Mars"); !ok{
		t.Logf("user1 got empty, but set to mars!")

		// got new set values
		if v, ok := syncMap.Load("user1"); ok {
			t.Logf("user1 exist =>%v", v)
		}
	}
}
