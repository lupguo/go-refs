package rdshand

import (
	"context"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/json-iterator/go"
	"github.com/pkg/errors"
)

// PipeWrite marshal objs to json string, then using pipeline set to redis
func (h *RdsHand) PipeWrite(ctx context.Context, objs []interface{}, kfn KeyFunc, t time.Duration) error {
	// make pipeline
	pipe := h.Rds.Pipeline()
	for _, obj := range objs {
		// serialize obj val
		val, err := jsoniter.MarshalToString(obj)
		if err != nil {
			continue
		}
		// set val
		pipe.Set(ctx, kfn(obj), val, t)
	}
	// exec pipeline
	_, err := pipe.Exec(ctx)
	if err != nil {
		return errors.Wrap(err, "PipeWrite() pipeline exec error")
	}
	return nil
}

// PipeRead using pipeline read multi keys to map[string]string for each element
func (h *RdsHand) PipeRead(ctx context.Context, keys []string) (map[string]string, error) {
	if len(keys) == 0 {
		return nil, errors.New("PipeRead() read empty keys")
	}
	// make pipeline
	pipe := h.Rds.Pipeline()
	mcmd := map[string]*redis.StringCmd{}
	for _, key := range keys {
		mcmd[key] = pipe.Get(ctx, key)
	}
	// exec pipeline
	_, err := pipe.Exec(ctx)
	if err != nil && err != redis.Nil {
		return nil, errors.Wrap(err, "PipeRead() pipeline exec error")
	}

	// read from pipeline
	out := map[string]string{}
	for k, v := range mcmd {
		if err := v.Err(); err != nil {
			if err == redis.Nil {
				continue
			}
			return nil, err
		}
		out[k] = v.Val()
	}

	return out, nil
}

// DelKeys using redis client del multi keys
func (h *RdsHand) DelKeys(ctx context.Context, keys ...string) error {
	if len(keys) == 0 {
		return nil
	}
	return h.Rds.Del(ctx, keys...).Err()
}
