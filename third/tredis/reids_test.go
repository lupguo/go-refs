package tredis

import (
	"context"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

/**
go redis包:
1. 实例创建:
	- 指定指定Options值，`redis.NewClient(&redis.Options{})`
	- 通过`opts, err := redis.ParseURL()`解析得到opts，再`rdb := redis.NewClient(opts)`
2. 请求发送:
	- `val, err := rdb.Get(ctx, "key").Result()`，判断err后，要断言返回的val值，直接通过附加组手方法获取值:`get.Int()`,`get.Text()`
	- `val, err := rdb.Do(ctx, "get", "key").Result()`
3. Error处理
	- redis.Nil判断，redis返回结果为nil转化成go的`redis.Nil`变量
	- 空值判断:
		- err == redis.Nil: 值未设置
		- err redis操作错误
		- val == "": 有key，但key的值为空
4. 类型转换:
	- 通过组手方法实现：`get.Text()`,`get.Int()`,`get.Uint64()`,`get.Float32()`,`get.Bool()`...
	-
*/

var (
	// step1
	opts = &redis.Options{
		Addr:       "127.0.0.1:6379",
		Password:   "clark", // no password set
		DB:         2,       // use default DB
		MaxRetries: 3,
	}
	rdb = redis.NewClient(opts)
	ctx = context.Background()
)

func init() {
	// step2
	// [scheme:][//[username[:password]@]host][/]path[?query][#fragment]
	opts, err := redis.ParseURL("redis://:clark@127.0.0.1:6379/2")
	if err != nil {
		panic(err)
	}
	rdb = redis.NewClient(opts)
}

// redis
func TestQuickStart(t *testing.T) {
	err := rdb.Set(ctx, "hey", "Hey, man!", 0).Err()
	if err != nil {
		panic(err)
	}
	val, err := rdb.Get(ctx, "hey").Result()
	if err == redis.Nil {
		t.Logf("get hey not exist, err: %+v", err)
	} else if err != nil {
		panic(err)
	}
	t.Logf("get hey vals: %s", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	// redis.Nil判断key不存在（类似于mysql的no found record）
	if err == redis.Nil {
		t.Log("key2 does not exist")
	} else if err != nil {
		panic(err)
	}
	t.Log("key2", val2)
}

func TestRedisGet(t *testing.T) {
	// get string
	get, err := rdb.Get(ctx, "hey").Result()
	assert.Nil(t, err)
	t.Logf("get hey return:%s", get)
	get, err = rdb.Get(ctx, "hey1").Result()
	assert.Equal(t, err, redis.Nil)
	t.Logf("get hey1 return nil")
}

func TestRedisDo(t *testing.T) {
	val, err := rdb.Do(ctx, "get", "hey").Result()
	if err != nil {
		if err == redis.Nil {
			t.Logf("key does not exists")
			return
		}
		panic(err)
	}
	t.Logf(val.(string))
}

// https://www.redis.com.cn/commands.html
func TestKeyValueLookAndFree(t *testing.T) {
	// SET key value EX 10 NX
	set, err := rdb.SetNX(ctx, "key", "value", 60*time.Second).Result()
	assert.Nil(t, err)
	t.Logf("setnx value: %t", set)

	// SET key value keepttl NX
	set, err = rdb.SetNX(ctx, "key", "value", redis.KeepTTL).Result()
	assert.Nil(t, err)
	t.Logf("set value: %t", set)

	ttl, err := rdb.TTL(ctx, "key").Result()
	assert.Nil(t, err)
	t.Logf("ttl left:%s", ttl)
}

// List
// Sort 数字、字符修饰符排序、外部排序
func TestRedisList(t *testing.T) {
	key := "list"
	list, err := rdb.RPush(ctx, key, []interface{}{1, 3, 4, 5, 7, 8, 9, 2, 0, 6}).Result()
	assert.Nil(t, err)
	t.Logf("rpush list return num:%d", list)

	// exist
	exist, err := rdb.Exists(ctx, key).Result()
	assert.Nil(t, err)
	t.Logf("exists key:%d", exist)

	// expire
	expBool, err := rdb.Expire(ctx, key, 80*time.Second).Result()
	assert.Nil(t, err)
	t.Logf("expire rs:%t", expBool)

	// ttl
	ttl, err := rdb.TTL(ctx, key).Result()
	assert.Nil(t, err)
	t.Logf("ttl expire 80:%v", ttl)

	time.Sleep(2 * time.Second)

	// ttl
	ttl, err = rdb.TTL(ctx, key).Result()
	assert.Nil(t, err)
	t.Logf("ttl expire 80:%v", ttl)

	// persis
	if _, err := rdb.Persist(ctx, key).Result(); err != nil {
		return
	}

	// ttl again
	ttl, err = rdb.TTL(ctx, key).Result()
	assert.Nil(t, err)
	t.Logf("ttl persis :%v", ttl)

	// SORT list LIMIT 0 2 ASC
	sortList, err := rdb.Sort(ctx, key, &redis.Sort{Offset: 0, Count: 15, Order: "ASC"}).Result()
	assert.Nil(t, err)
	t.Logf("sort list values: %+v", sortList)
}

func TestRedisZsets(t *testing.T) {
	// ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
	zrVals, err := rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  2,
	}).Result()
	assert.Nil(t, err)
	t.Logf("zrangesoce %v", zrVals)

	// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
	ziVals, err := rdb.ZInterStore(ctx, "out", &redis.ZStore{
		Keys:    []string{"zset1", "zset2"},
		Weights: []float64{2, 3},
	}).Result()
	assert.Nil(t, err)
	t.Logf("ziVals:%d", ziVals)
}

func TestRedisEval(t *testing.T) {
	// EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
	ziVals, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()
	assert.Nil(t, err)
	t.Logf("eval vaules:%+v", ziVals)

	// custom command
	res, err := rdb.Do(ctx, "set", "key", "value").Result()
	assert.Nil(t, err)
	t.Logf("set return: %+v", res)
}
