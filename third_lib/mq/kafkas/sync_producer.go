package kafkas

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

var (
	syncProducerOnce sync.Once
	syncProducerMap  map[string]*SyncProducer
)

// SyncProducer sarama 生产者
type SyncProducer struct {
	BrokersUrl   []string
	Config       *sarama.Config
	syncProducer sarama.SyncProducer
}

// NewSyncProducer 初始化同步生产者实例。如果之前已经创建过，直接获取用于后续MQ消息发送；否则利用sync.once完成初始化
func NewSyncProducer(ctx context.Context, name string, brokersUrl []string, cfg *sarama.Config) (producer *SyncProducer, err error) {
	// 已初始化，直接返回
	if syncProducerMap[name] != nil {
		return syncProducerMap[name], nil
	}

	// once初始化
	syncProducerOnce.Do(func() {
		producer, err = newSyncProducer(ctx, brokersUrl, cfg)
		if err != nil {
			err = errors.Wrap(err, "kafkas once.Do() new sync producer got err")
			return
		}
		syncProducerMap = make(map[string]*SyncProducer)
		syncProducerMap[name] = producer
	})

	return syncProducerMap[name], nil
}

// newSyncProducer 创建一个新的Producer生产者，用于后续消息发布，cfg 传递为空，则采用默认配置
func newSyncProducer(ctx context.Context, brokersUrl []string, cfg *sarama.Config) (*SyncProducer, error) {
	// 资源poll没有拿到，重新创建(需要考虑短连接过多的问题)
	if cfg == nil {
		cfg = sarama.NewConfig()
		cfg.Producer.Retry.Max = 3
		cfg.Producer.RequiredAcks = sarama.WaitForAll
		cfg.Producer.Return.Successes = true
	}

	// 初始化同步生产者
	syncProducer, err := sarama.NewSyncProducer(brokersUrl, cfg)
	if err != nil {
		return nil, errors.Wrap(err, "kafkas sarama new sync producer got err")
	}

	return &SyncProducer{
		BrokersUrl:   brokersUrl,
		Config:       cfg,
		syncProducer: syncProducer,
	}, nil
}

// SyncPushMessageToQueue 推送消息到消息队列
func (p *SyncProducer) SyncPushMessageToQueue(ctx context.Context, topic string, message []byte) (err error) {
	if p.syncProducer == nil {
		return errors.New("p.syncProducer got error, p.syncProducer is nil")
	}
	msg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
	}

	if _, _, err = p.syncProducer.SendMessage(msg); err != nil {
		return errors.Wrap(err, "kafka send message got err")
	}

	return nil
}
