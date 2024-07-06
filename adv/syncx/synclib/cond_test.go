package synclib

import (
	"math/rand"
	"sync"
	"testing"
	"time"
)

func TestCond(t *testing.T) {
	rand.Seed(time.Now().UnixNano())
	rwLocker := &sync.RWMutex{}
	syncCond := sync.NewCond(rwLocker)
	syncCond.L.Lock()
	go func() {
		for {
			select {
			case <-time.Tick(time.Duration(rand.Int63n(3)) * time.Second):
				t.Logf("case1: exec..")
			case <-time.Tick(time.Duration(rand.Int63n(3)) * time.Second):
				t.Logf("case2: Now time [%s]\n", time.Now().Format("2006-01-02 15:04:05"))
			case <-time.Tick(2 * time.Second):
				t.Logf("case3: exec..")
				syncCond.Signal()
			}
		}
	}()
	syncCond.Wait()
}
