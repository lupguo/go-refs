package verifycode

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestRedisVerifier_TokenVerify(t *testing.T) {
	var opts *redis.Options
	opts, err := redis.ParseURL("redis://:clark@127.0.0.1:6379/0")
	assert.Nil(t, err)

	// 生成token
	vry := NewRedisVerifier(redis.NewClient(opts), "salt", "token:")
	ctx := context.Background()
	token, err := vry.MakeToken(ctx, 1851234567, 60*time.Second)
	assert.Nil(t, err)

	// 校验token
	val, err := vry.VerifyToken(ctx, token)
	assert.Nil(t, err)

	t.Logf("vaule return: %v", val)
}

func BenchmarkRedisVerifier_MakeToken(b *testing.B) {
	var opts *redis.Options
	opts, err := redis.ParseURL("redis://:clark@127.0.0.1:6379/0")
	assert.Nil(b, err)

	for i := 0; i < b.N; i++ {
		// 生成token
		vry := NewRedisVerifier(redis.NewClient(opts), "salt", "token:")
		ctx := context.Background()
		_, err = vry.MakeToken(ctx, 1851234567, 60*time.Second)
		assert.Nil(b, err)
	}
}