package main

import (
	"fmt"
	"os"
	"runtime"
	"runtime/pprof"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/spf13/pflag"
)

var (
	address string
	topic   string
	data    string
)

func main() {
	runtime.GOMAXPROCS(8)

	memProfile, _ := os.Create("/tmp/mem_profile")
	cpuProfile, _ := os.Create("/tmp/cpu_profile")
	pprof.StartCPUProfile(cpuProfile)
	pprof.WriteHeapProfile(memProfile)
	defer pprof.StopCPUProfile()

	pflag.StringVarP(&address, "address", "a", "kafka_dev_node:9092", "kafka address")
	pflag.StringVarP(&topic, "topic", "t", "account_event", "kafka topic")
	pflag.StringVarP(&data, "data", "d", `{key: "val"}`, "send data")
	pflag.Parse()

	// convert body into bytes and send it to kafka
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	// config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 2
	asyncProducer, err := sarama.NewAsyncProducer([]string{address}, config)
	if err != nil {
		panic(err)
	}

	// concurrent send message to mq
	ConcurrentSyncSendMessage(asyncProducer)
}

// ConcurrentSyncSendMessage multi push message
func ConcurrentSyncSendMessage(producer sarama.AsyncProducer) {
	// send
	go func() {
		sema := make(chan bool, 8)
		wg := sync.WaitGroup{}
		for i := 0; i < 1e5; i++ {
			wg.Add(1)
			sema <- true
			go func(i int) {
				defer wg.Done()

				// msg
				msg := &sarama.ProducerMessage{
					Topic: topic,
					Value: sarama.StringEncoder(fmt.Sprintf("{id:%04d}", i)),
				}

				producer.Input() <- msg
			}(i)
			<-sema
		}
		wg.Wait()
		producer.AsyncClose()
	}()

	// handle result
	wg := sync.WaitGroup{}
	wg.Add(1)
	var cntOk, cntNok int
	go func() {
		defer wg.Done()
		for succ := range producer.Successes() {
			cntOk++
			fmt.Printf("succ, %s\n", succ.Value)
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		for err := range producer.Errors() {
			cntNok++
			fmt.Printf("err, %+v\n", err)
		}
	}()

	wg.Wait()

	// last
	fmt.Printf("cntOk=%d, cntNok=%d", cntOk, cntNok)
}
