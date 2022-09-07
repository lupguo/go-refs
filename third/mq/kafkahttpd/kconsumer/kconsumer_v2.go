package kconsumer

import (
	"context"
	"log"

	"github.com/segmentio/kafka-go"
)

type KConsumer struct {
	brokersURLs []string
	readCfg     kafka.ReaderConfig
}

func NewKConsumer(brokersURLs []string) *KConsumer {
	return &KConsumer{
		brokersURLs: brokersURLs,
	}
}

var (
	readerMinBytes int = 10e2
	readerMaxBytes int = 10e4
)

// createMessageReader 创建消息读取对象
func (c *KConsumer) createMessageReader(topic string, groupID string) *kafka.Reader {
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  c.brokersURLs,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: readerMinBytes, // 10KB
		MaxBytes: readerMaxBytes, // 10MB
	})
}

// ConsumeMessageV2 消息消费
func (c *KConsumer) ConsumeMessageV2(ctx context.Context, topic string, groupID string) error {
	reader := c.createMessageReader(topic, groupID)
	defer reader.Close()

	for {
		m, err := reader.ReadMessage(ctx)
		if err != nil {
			panic(err)
		}
		log.Printf("message at topic:%v | partition:%v |offset:%v | message key => %s , message value => %s\n", m.Topic, m.Partition, m.Offset, string(m.Key), string(m.Value))
	}
}
