package kafkas

import (
	"context"
	"fmt"
	"os"
	"runtime/pprof"
	"sync"
	"testing"
	"time"
)

// 	- 多协程:
//		BenchmarkSyncProducer-12    	     222	   7252598 ns/op(50ms耗时) - no poll 不复用
//		BenchmarkSyncProducer-12    	     550	   3573820 ns/op(50ms耗时) - poll 网络连接复用，调用平次越高复用度越大
func BenchmarkSyncProducer(b *testing.B) {
	b.SkipNow()

	memProfile, _ := os.Create("/tmp/mem_profile")
	cpuProfile, _ := os.Create("/tmp/cpu_profile")
	pprof.StartCPUProfile(cpuProfile)
	pprof.WriteHeapProfile(memProfile)
	defer pprof.StopCPUProfile()

	brokerURLs := []string{
		"kafka_dev_node:9092",
	}

	ctx := context.Background()
	wg := sync.WaitGroup{}
	name := `benchT`
	// sema := make(chan bool, 10)
	for i := 0; i < 1e4; i++ {
		wg.Add(1)
		go func() {
			defer func() {
				wg.Done()
			}()
			producer, err := NewSyncProducer(ctx, name, brokerURLs, nil)
			if err != nil {
				b.Error(err)
			}
			err = producer.SyncPushMessageToQueue(context.Background(), "my-topic", []byte(fmt.Sprintf(`{"time": "%s"}`, time.Now().Format("2006/01/02 15:04:05"))))
			if err != nil {
				b.Logf("got err: %s", err)
			}
			// 模拟50ms耗时
			time.Sleep(100 * time.Millisecond)
		}()
	}
	wg.Wait()
}
