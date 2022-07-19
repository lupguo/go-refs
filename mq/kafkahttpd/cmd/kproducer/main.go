package main

import (
	"log"
	"sync"

	"github.com/spf13/pflag"
	"x-learn/mq/kafkahttpd/kproducer"
)

var (
	address string
	topic   string
	data    string
)

func main() {
	pflag.StringVarP(&address, "address", "a", "kafka_dev_node:9092", "kafka address")
	pflag.StringVarP(&topic, "topic", "t", "account_event", "kafka topic")
	pflag.StringVarP(&data, "data", "d", `{key: "val"}`, "send data")
	pflag.Parse()

	// convert body into bytes and send it to kafka
	producer := kproducer.NewKProducer([]string{address})

	// concurrent send message to mq
	ConcurrentSyncSendMessage(producer)
}

// ConcurrentSyncSendMessage multi push message
func ConcurrentSyncSendMessage(producer *kproducer.KProducer) {
	wg := sync.WaitGroup{}
	for i := 0; i < 1e3; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			err := producer.PushMessageToQueue(topic, []byte(data))
			if err != nil {
				log.Fatalln("push got err:", err)
			}
		}()
	}
	wg.Wait()
}


