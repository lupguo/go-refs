package redismock

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/go-redis/redismock/v8"
	"github.com/pkg/errors"
	"github.com/stretchr/testify/assert"
)

var ctx = context.TODO()

func NewsInfoForCache(redisDB *redis.Client, newsID int) (info string, err error) {
	cacheKey := fmt.Sprintf("news_redis_cache_%d", newsID)
	info, err = redisDB.Get(ctx, cacheKey).Result() // 触发Get Expect，从redis取值，如果为Nil值，则直接做Set
	if err == redis.Nil {
		// info, err = call api()
		info = "test"
		err = redisDB.Set(ctx, cacheKey, info, 30*time.Minute).Err()	// 触发Set Expect
	}
	return
}

func TestNewsInfoForCache(t *testing.T) {
	db, mock := redismock.NewClientMock()

	newsID1 := 123456789
	key1 := fmt.Sprintf("news_redis_cache_%d", newsID1)
	newsID2 := 110
	key2 := fmt.Sprintf("news_redis_cache_%d", newsID2)

	// mock ignoring `call api()`
	// expect期望设置
	mock.ExpectGet(key1).RedisNil()
	mock.Regexp().ExpectSet(key1, `[a-z]+`, 30 * time.Minute).SetErr(errors.New("FAIL"))

	// test get info from cache
	_, err := NewsInfoForCache(db, newsID1)
	if err == nil || err.Error() != "FAIL" {
		t.Error("wrong error")
	}

	// expect期望设置
	mock.ExpectGet(key2).SetVal("Exist")
	info, err2 := NewsInfoForCache(db, newsID2)
	assert.Equal(t, info, "Exist")
	assert.Nil(t, err2)

	// ?
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Error(err)
	}
}