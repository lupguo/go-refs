package goretryx

import (
	"context"
	"testing"
	"time"

	"github.com/avast/retry-go/v4"
	"github.com/pkg/errors"
)

// goretryx_test.go:14: [2022-09-05 16:39:54.965643 +0800 CST m=+0.000860792] test start
// goretryx_test.go:36: [2022-09-05 16:39:54.965876 +0800 CST m=+0.001094417], cost time: 233.709µs, retry #0, because got err: uid 100 err
// goretryx_test.go:36: [2022-09-05 16:39:55.114179 +0800 CST m=+0.149396459], cost time: 148.535709ms, retry #1, because got err: uid 101 err
// goretryx_test.go:36: [2022-09-05 16:39:55.3974 +0800 CST m=+0.432616834], cost time: 431.756084ms, retry #2, because got err: uid 102 err
// goretryx_test.go:54: [2022-09-05 16:39:55.69802 +0800 CST m=+0.733236501] exec ok
func TestRetry(t *testing.T) {
	start := time.Now()
	t.Logf("[%s] test start", start)

	// 设定超时context
	ctx := context.Background()
	// ctx, cancel := context.WithTimeout(ctx, 250*time.Millisecond)
	go func() {
		// 若cancel执行，后续retry重试捕获到则直接退出
		// select {
		// case <-time.After(50 * time.Millisecond):
		// 	cancel()
		// }
		// _ = cancel
	}()

	// retry配置
	retryOpts := []retry.Option{
		retry.Context(ctx),                     // 支持传入带超时限定的ctx，如果ctx超时，但还在重试过程中，则直接返回ctx超时错误 基于select case实现）
		retry.Attempts(4),                      // 尝试重试的次数
		retry.Delay(200 * time.Millisecond),    // Delay结合MaxDelay、以及退避算法综合得到 （基于select case实现）
		retry.MaxDelay(500 * time.Millisecond), // 重试最大间隔时间，如果退避后的间隔时间超过maxDelay，则改用maxDelay
		retry.LastErrorOnly(true),              // 仅返回最后一次错误
		retry.OnRetry(func(n uint, err error) { // 每次重试的时候调用方法
			t.Logf("[%s], cost time: %s, retry #%d, because got err: %s", time.Now(), time.Now().Sub(start), n, err)
		}),
		retry.DelayType(retry.CombineDelay(retry.BackOffDelay, retry.RandomDelay)), // 这个就是默认组合: 退避+随机抖动延迟(maxJitter)
	}

	uid := 100
	err := retry.Do(func() error {
		switch uid {
		case 100:
			uid++
			return errors.New("uid 100 err")
		case 101:
			uid++
			return errors.New("uid 101 err")
		case 102:
			uid++
			return errors.New("uid 102 err")
		}
		t.Logf("[%s] exec ok", time.Now().String())
		return nil
	}, retryOpts...)

	if err != nil {
		t.Error("exec got err:", err)
	}
}
