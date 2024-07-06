package kafkas

import (
	"context"
	"sync"

	"git.code.oa.com/trpc-go/trpc-go/log"
	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

var (
	asyncProducerOnce sync.Once
	asyncProducerMap  map[string]*AsyncProducer
)

// AsyncProducer 异步生产者
type AsyncProducer struct {
	BrokersUrl    []string
	Config        *sarama.Config
	asyncProducer sarama.AsyncProducer
}

// NewASyncProducer 创建一个新的Producer生产者，用于后续消息发布
// 	cfg 传递为空，则采用默认配置
func NewASyncProducer(ctx context.Context, name string, brokersUrl []string, cfg *sarama.Config) (*AsyncProducer, error) {
	// 已初始化，直接返回
	if asyncProducerMap[name] != nil {
		return asyncProducerMap[name], nil
	}

	// once初始化
	asyncProducerOnce.Do(func() {
		producer, err := newAsyncProducer(ctx, brokersUrl, cfg)
		if err != nil {
			log.ErrorContextf(ctx, "kafkas once.Do() new async producer got err: %s", err)
			return
		}
		asyncProducerMap = make(map[string]*AsyncProducer)
		asyncProducerMap[name] = producer
	})

	return asyncProducerMap[name], nil
}

// newAsyncProducer 创建一个异步的Producer生产者
func newAsyncProducer(ctx context.Context, brokersUrl []string, cfg *sarama.Config) (*AsyncProducer, error) {
	if cfg == nil {
		cfg = sarama.NewConfig()
		cfg.Producer.Retry.Max = 3
		cfg.Producer.Return.Successes = true
	}

	// 初始化同步生产者
	asyncProducer, err := sarama.NewAsyncProducer(brokersUrl, cfg)
	if err != nil {
		log.ErrorContextf(ctx, "kafka new async producer got err: %v", err)
		return nil, err
	}

	// process async log
	go func() {
		for {
			select {
			case err := <-asyncProducer.Errors():
				log.ErrorContextf(ctx, "p.asyncProducer AsyncPushMessageToQueue() got err, %v", err)
			case msg := <-asyncProducer.Successes():
				log.DebugContextf(ctx, "message is stored in topic(%s)/partition(%d)/offset(%d)", msg.Topic, msg.Partition, msg.Offset)
			}
		}
	}()

	return &AsyncProducer{
		BrokersUrl:    brokersUrl,
		Config:        cfg,
		asyncProducer: asyncProducer,
	}, nil
}

// AsyncPushMessageToQueue 异步投递消息到Kafka
func (p *AsyncProducer) AsyncPushMessageToQueue(ctx context.Context, topic string, message []byte) (err error) {
	// 异步消息发送
	if p.asyncProducer == nil {
		return errors.New("p.asyncProducer is nil, can not async push message to queue")
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	// 异步消息投递
	p.asyncProducer.Input() <- msg
	log.DebugContextf(ctx, "message async push done, msg: %s", message)

	return nil
}
