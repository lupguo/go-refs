package curr_lock

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestDoThingsV1(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			time.Sleep(time.Duration(rand.Int63n(10))*time.Microsecond)
			t.Logf("handle %d, got %v", id, DoThingsV1())
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func TestDoThingsV2(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			time.Sleep(time.Duration(rand.Int63n(10))*time.Microsecond)
			t.Logf("handle %d, got %v", id, DoThingsV2())
			wg.Done()
		}(i)
	}
	wg.Wait()
}
