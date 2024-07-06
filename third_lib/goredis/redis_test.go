package goredis

import (
	"context"
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func getRds() *redis.Client {
	rdsOpts, err := redis.ParseURL(`redis://clark:Secret@123.@9.134.233.187:6380/0`)
	if err != nil {
		log.Fatalf("getRds() fail: %s", err)
	}

	return redis.NewClient(rdsOpts)
}

var (
	ctx = context.Background()
)

func TestRedisGetCommand(t *testing.T) {
	rds := getRds()
	// return ""
	v1 := rds.Get(ctx, `tst_100`).Val()
	assert.Equal(t, v1, "")
	t.Logf("v1:%+v", v1)

	// return redis: nil
	vs := rds.Get(ctx, `tst_100`).String()
	assert.Equal(t, vs, "get tst_100: redis: nil")
	t.Logf("vs:%+v", vs)

	assert.NotEqual(t, v1, vs)
}

func TestExampleNewClient(t *testing.T) {
	rds := getRds()

	pong, err := rds.Ping(ctx).Result()
	t.Log(pong, err)
	// Output: PONG <nil>
}

func TestExampleClient(t *testing.T) {
	rdb := getRds()
	err := rdb.Set(ctx, "key", "value", 0).Err()
	if err != nil {
		t.Fatal(err)
	}

	val, err := rdb.Get(ctx, "key").Result()
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("key", val)

	val2, err := rdb.Get(ctx, "key2").Result()
	if err == redis.Nil {
		fmt.Println("key2 does not exist")
	} else if err != nil {
		t.Fatal(err)
	} else {
		t.Log("key2", val2)
	}
	// Output: key value
	// key2 does not exist
}

func TestExists(t *testing.T) {
	rds := getRds()

	// not exist key
	existKey3 := rds.Exists(ctx, "key3")
	cmd, err := existKey3.Result()
	t.Logf("cmd=%+v, err=%v, val=%v, str=%v", cmd, err, existKey3.Val(), existKey3.String())

	// exist key
	rds.Set(ctx, "hello", "val", -1)
	existHello := rds.Exists(ctx, "hello")
	cmd, err = existHello.Result()
	t.Logf("cmd=%+v, err=%v, val=%v, str=%v", cmd, err, existHello.Val(), existHello.String())
}

func TestExpire(t *testing.T) {
	rds := getRds()
	key := "expireKey"
	// del
	t.Logf("del err=%v", rds.Del(ctx, key).Err())

	// set
	err := rds.Set(ctx, key, time.Now().String(), -1).Err()
	t.Logf("err=%v", err)

	// get
	t.Logf("exist val=%v", rds.Exists(ctx, key).Val())

	// ttl
	t.Logf("ttl=%v", rds.TTL(ctx, key).Val())

}

//    redis_test.go:91: true true <nil> <nil> ok
//    redis_test.go:99: false false <nil> <nil> ok
func TestSetNx(t *testing.T) {
	rdsOpts, err := redis.ParseURL("redis://:clark@127.0.0.1:6379/0")
	assert.Nil(t, err)
	rdb := redis.NewClient(rdsOpts)

	// setnx 1th
	doSetnx := rdb.SetNX(ctx, "key", "ok", 10*time.Second)
	snxVal := doSetnx.Val()
	snxRst, rstErr := doSetnx.Result()
	snxErr := doSetnx.Err()
	// 1th val
	snxVV := rdb.Get(ctx, "key").Val()
	t.Log(snxVal, snxRst, rstErr, snxErr, snxVV)

	// 2th
	doSetnx = rdb.SetNX(ctx, "key", "ok", 10*time.Second)
	snxVal = doSetnx.Val()
	snxRst, rstErr = doSetnx.Result()
	snxErr = doSetnx.Err()
	snxVV = rdb.Get(ctx, "key").Val()
	t.Log(snxVal, snxRst, rstErr, snxErr, snxVV)
}

// [scheme:][//[userinfo@]host][/]path[?query][#fragment]
func TestGoRedis(t *testing.T) {
	rdsOpts, err := redis.ParseURL("redis://:clark@127.0.0.1:6379/0")
	assert.Nil(t, err)

	rdb := redis.NewClient(rdsOpts)

	// SET key value EX 10 NX
	set, err := rdb.SetNX(ctx, "key", "value", 10*time.Second).Result()
	assert.Nil(t, err)
	t.Log(set)

	// SET key value keepttl NX
	set, err = rdb.SetNX(ctx, "key", "value", redis.KeepTTL).Result()
	assert.Nil(t, err)
	t.Log(set)

	// SORT list LIMIT 0 2 ASC
	vals, err := rdb.Sort(ctx, "list", &redis.Sort{Offset: 0, Count: 2, Order: "ASC"}).Result()
	assert.Nil(t, err)
	t.Log(vals)

	// ZRANGEBYSCORE zset -inf +inf WITHSCORES LIMIT 0 2
	zvals, err := rdb.ZRangeByScoreWithScores(ctx, "zset", &redis.ZRangeBy{
		Min:    "-inf",
		Max:    "+inf",
		Offset: 0,
		Count:  2,
	}).Result()
	assert.Nil(t, err)
	t.Logf("zrange:%+v", zvals)

	// ZINTERSTORE out 2 zset1 zset2 WEIGHTS 2 3 AGGREGATE SUM
	zivals, err := rdb.ZInterStore(ctx, "out", &redis.ZStore{
		Keys:    []string{"zset1", "zset2"},
		Weights: []float64{2.0, 3.0},
	}).Result()
	assert.Nil(t, err)
	t.Logf("zival:%+v", zivals)

	// EVAL "return {KEYS[1],ARGV[1]}" 1 "key" "hello"
	evalVals, err := rdb.Eval(ctx, "return {KEYS[1],ARGV[1]}", []string{"key"}, "hello").Result()
	assert.Nil(t, err)
	t.Logf("eval %+v", evalVals)

	// custom command
	res, err := rdb.Do(ctx, "set", "key", "value").Result()
	assert.Nil(t, err)
	t.Logf("custom command: %+v", res)
}

func TestInitBdata(t *testing.T) {
	client := getRds()
	for i := 0; i < 10000; i++ {
		client.Set(ctx, "key"+strconv.Itoa(i), "hoge"+strconv.Itoa(i), redis.KeepTTL)
	}
	// client := redis.NewClient(&redis.Options{Addr: s.Addr()})
}

func noUsingPipeline() {
	// 普通，多次网络开销
	client := getRds()
	result := map[string]string{}
	for i := 0; i < 10000; i++ {
		key := "key" + strconv.Itoa(i)
		res, _ := client.Get(ctx, key).Result()
		result[key] = res
	}
}

func usingPipeline() {
	client := getRds()
	// Pipeline，一次网络开销
	m := map[string]*redis.StringCmd{}
	pipe := client.Pipeline()
	for i := 0; i < 10000; i++ {
		m["key"+strconv.Itoa(i)] = pipe.Get(ctx, "key"+strconv.Itoa(i))
	}
	_, err := pipe.Exec(ctx)
	if err != nil {
		panic(err)
	}

	result2 := map[string]string{}
	for k, v := range m {
		res, _ := v.Result()
		result2[k] = res
	}
}

// goredis_test.go:155: no using pipeline elpase 8.820800906s
// goredis_test.go:158: using pipeline elpase  160.278788ms
func TestPipeline(t *testing.T) {
	t1 := time.Now()
	noUsingPipeline()
	t.Logf("no using pipeline elpase %s", time.Since(t1))

	t2 := time.Now()
	usingPipeline()
	t.Logf("using pipeline elpase  %s", time.Since(t2))
}
