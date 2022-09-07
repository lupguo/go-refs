package kconsumer

import (
	"context"
	"log"

	"github.com/Shopify/sarama"
)

type PartitionConsumer struct {
	BrokersUrl []string
}

func NewPartitionConsumer(brokersUrl []string) *PartitionConsumer {
	return &PartitionConsumer{BrokersUrl: brokersUrl}
}

// connectConsumer 连接消费者brokers
func (pc *PartitionConsumer) connectConsumer() (sarama.Consumer, error) {
	config := sarama.NewConfig()
	config.Consumer.Return.Errors = true

	// client, err := sarama.NewConsumerGroup(strings.Split(brokersURLs, ","), group, config)
	// NewConsumer creates a new consumer using the given broker addresses and configuration
	consumer, err := sarama.NewConsumer(pc.BrokersUrl, config)
	if err != nil {
		return nil, err
	}
	return consumer, nil
}

// ConsumeMessage 消息消费
func (pc *PartitionConsumer) ConsumeMessage(ctx context.Context, partition int32, topic string) error {
	consumer, err := pc.connectConsumer()
	if err != nil {
		panic(err)
	}
	// calling ConsumePartition. It will open one connection per broker
	// and share it for all partitions that live on it.
	pconsumer, err := consumer.ConsumePartition(topic, partition, sarama.OffsetNewest)
	defer pconsumer.Close()

	var count = 0
	for {
		select {
		case msg := <-pconsumer.Messages():
			count++
			log.Printf("MsgCount %d: | Partition(%d) | Topic(%s) | Message(%s) | Offset(%d)", count, msg.Partition, msg.Topic, msg.Value, msg.Offset)
		case err := <-pconsumer.Errors():
			log.Printf("Got Error %s", err)
		}
	}
}
