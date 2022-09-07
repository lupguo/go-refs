package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/google/uuid"
	"github.com/segmentio/kafka-go"
)

func newKafkaWriter(kafkaURL, topic string) *kafka.Writer {
	return &kafka.Writer{
		Addr:     kafka.TCP(kafkaURL),
		Topic:    topic,
		Balancer: &kafka.LeastBytes{},
	}
}

var (
	brokersUrl = "kafka_dev_node:9092"
	ctx        = context.Background()
	topic      = "comments"
)

// 消息生产
func main() {
	for i := 0; ; i++ {
		// 消息key和内容
		key := fmt.Sprintf("Key-%d", i)
		msg := kafka.Message{
			Key:   []byte(key),
			Value: []byte(fmt.Sprint(uuid.New())),
		}

		// 新建kafka写
		writer := newKafkaWriter(brokersUrl, topic)
		defer writer.Close()

		fmt.Println("start producing ... !!")
		err := writer.WriteMessages(ctx, msg)
		if err != nil {
			log.Println(err)
			continue
		}
		fmt.Println("produced", key)
		time.Sleep(5 * time.Millisecond)
	}
}
