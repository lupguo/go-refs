package kproducer

import (
	"fmt"

	"github.com/Shopify/sarama"
)

type KProducer struct {
	BrokersUrl   []string
	syncProducer sarama.SyncProducer
}

func NewKProducer(brokersUrl []string) *KProducer {
	return &KProducer{
		BrokersUrl:   brokersUrl,
	}
}

// connectProducer 连接kafka服务，返回一个同步生产者
func (p *KProducer) connectProducer() (producer sarama.SyncProducer, err error) {
	if p.syncProducer != nil {
		return p.syncProducer, nil
	}

	// NewSyncProducer creates a new syncProducer using the given broker addresses and configuration.
	config := sarama.NewConfig()
	config.Producer.Return.Successes = true
	config.Producer.RequiredAcks = sarama.WaitForAll
	config.Producer.Retry.Max = 5
	p.syncProducer, err = sarama.NewSyncProducer(p.BrokersUrl, config)
	if err != nil {
		return nil, err
	}

	return p.syncProducer, nil
}

// PushMessageToQueue 推送消息到topick
func (p *KProducer) PushMessageToQueue(topic string, message []byte) error {
	producer, err := p.connectProducer()
	if err != nil {
		return err
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}
	partition, offset, err := producer.SendMessage(msg)
	if err != nil {
		return err
	}
	fmt.Printf("Message is stored in topic(%s)/partition(%d)/offset(%d)\n", topic, partition, offset)
	return nil
}
