package maps

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

// fatal error: concurrent map iteration and map write
func TestBenchWriteAndRangeRead(t *testing.T) {
	type nMap map[int]int

	mOld := make(map[int]int)
	mNew := make(map[int]int)

	ch := make(chan int, 1)

	// read
	go func(m nMap) {
		for k, v := range m {
			mNew[k] = v
		}
		ch <- 1
	}(mOld)

	// write
	go func(m nMap) {
		for i := 0; i < 1e5; i++ {
			m[i] = i
		}
	}(mOld)

	<-ch
	fmt.Printf("len of mOld=%d, len of mNew=%d", len(mOld), len(mNew))
}

// 通过sync map，限定并发写，确定range是否会动态range出来
// sync.map不会阻塞，range是基于当时的数值来的
func TestBenchSynMap(t *testing.T) {
	mOld := &sync.Map{}
	mNew := &sync.Map{}

	// count
	go func(m *sync.Map) {
		for {
			select {
			case <-time.Tick(1 * time.Second):
				t.Logf("tick length -->%v", lenSyncMap(m))
			}
		}
	}(mNew)

	// range old -> read+print
	go func(m *sync.Map) {
		for {
			select {
			case <-time.Tick(100 * time.Millisecond):
				m.Range(func(key, value interface{}) bool {
					mNew.Store(key, value)
					return true
				})
			}
		}
	}(mOld)

	// write old routine
	go func(m *sync.Map) {
		for i := 0; i < 1e2; i++ {
			time.Sleep(100 * time.Millisecond)
			m.Store(i, i)
		}
	}(mOld)
	time.Sleep(3 * time.Second)

	// -- finish
	fmt.Printf("len of mOld=%d, len of mNew=%d", lenSyncMap(mOld), lenSyncMap(mNew))
}

func lenSyncMap(smap *sync.Map) int {
	var length int
	smap.Range(func(key, value interface{}) bool {
		length++
		return true
	})
	return length
}
