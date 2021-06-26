package verifycode

import (
	"context"
	"crypto/md5"
	"fmt"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/pkg/errors"
	uuid "github.com/satori/go.uuid"
)

var ErrVerifyCodeNotFound = errors.New("verify code cannot be found")

// Verifier 校验器，支持token生成和校验
type Verifier interface {
	MakeToken(ctx context.Context, val interface{}, exp time.Duration) (string, error)
	VerifyToken(ctx context.Context, token string) (val interface{}, err error)
	genToken() string
}

// RedisVerifier redis校验器
type RedisVerifier struct {
	salt   string        // token生成时候加的盐值
	prefix string        // redis的key存储前缀，以:结尾
	rdb    *redis.Client // 以redis作为存储
}

// NewRedisVerifier 创建RedisVerify
func NewRedisVerifier(rdb *redis.Client, salt, prefix string) *RedisVerifier {
	return &RedisVerifier{
		rdb:    rdb,
		salt:   salt,
		prefix: prefix,
	}
}

// MakeToken 生成随机字符串，与val进行关联绑定，存储到Redis中
func (vry *RedisVerifier) MakeToken(ctx context.Context, val interface{}, exp time.Duration) (string, error) {
	// 生成换绑token
	token := vry.genToken()
	_, err := vry.rdb.SetNX(ctx, tokenRdbKey(vry.prefix, token), val, exp).Result()
	if err != nil {
		return "", errors.Wrapf(err, "update verify token setNx() got err")
	}
	return token, nil
}

// VerifyToken 校验换绑传入的token
func (vry *RedisVerifier) VerifyToken(ctx context.Context, token string) (interface{}, error) {
	// 读取redis
	val, err := vry.rdb.Get(ctx, tokenRdbKey(vry.prefix, token)).Result()
	if err != nil {
		if err == redis.Nil {
			return nil, ErrVerifyCodeNotFound
		}
		return nil, errors.Wrapf(err, "failed to get verify token (%s)", token)
	}
	return val, nil
}

// GetToken 获取认证token, 
func (vry *RedisVerifier) genToken() string {
	randStr := fmt.Sprintf("%s_%s", vry.salt, uuid.NewV4().String())
	return fmt.Sprintf("%x", md5.Sum([]byte(randStr)))
}

// tokenRdbKey redis中存储val值的key拼接
func tokenRdbKey(prefix string, token string) string {
	key := fmt.Sprintf("%s%s", prefix, token)
	return key
}
