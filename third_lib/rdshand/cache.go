package rdshand

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
)

// IRedisHand redis interface
type IRedisHand interface {
	GetClient() *redis.Client
	PipeWrite(ctx context.Context, objs []interface{}, kfn KeyFunc, t time.Duration) error
	PipeRead(ctx context.Context, keys []string) (map[string]string, error)
	DelKeys(ctx context.Context, keys ...string) error
}

// RdsHand RedisCache
type RdsHand struct {
	Rds *redis.Client
}

// New create redis obj
func New(client *redis.Client) *RdsHand {
	if client == nil {
		panic(client)
	}
	return &RdsHand{Rds: client}
}

// GetClient get client from RdsHand obj
func (h *RdsHand) GetClient() *redis.Client {
	return h.Rds
}
