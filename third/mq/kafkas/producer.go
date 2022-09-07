package kafkas

import (
	"context"
)

type MessageProducer interface {
	// SyncPushMessageToQueue 同步消息推送
	SyncPushMessageToQueue(ctx context.Context, topic string, message []byte) error

	// AsyncPushMessageToQueue 异步消息推送
	AsyncPushMessageToQueue(ctx context.Context, topic string, message []byte) error
}

// Producer 未用到协程池
type Producer struct {
	asyncProducer *AsyncProducer
	syncProducer  *SyncProducer
}

// NewProducer 新建一个生产集合
func NewProducer(asyncProducer *AsyncProducer, syncProducer *SyncProducer) *Producer {
	return &Producer{
		asyncProducer: asyncProducer,
		syncProducer:  syncProducer,
	}
}

// SyncPushMessageToQueue 同步推送消息到消息队列
func (p *Producer) SyncPushMessageToQueue(ctx context.Context, topic string, message []byte) error {
	return p.syncProducer.SyncPushMessageToQueue(ctx, topic, message)
}

// AsyncPushMessageToQueue 异常推送消息到消息队列
func (p *Producer) AsyncPushMessageToQueue(ctx context.Context, topic string, message []byte) error {
	return p.asyncProducer.AsyncPushMessageToQueue(ctx, topic, message)
}
