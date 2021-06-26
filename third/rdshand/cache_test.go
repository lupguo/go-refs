package rdshand

import (
	"context"
	"encoding/json"
	"testing"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/stretchr/testify/assert"
)

func TestSimpleSet(t *testing.T) {
	hand := New(getRds())
	rds := hand.GetClient()
	key := RdKey(`simple:%d`, time.Now().Unix())
	_, err := rds.Get(ctx, key).Bool()
	assert.Equal(t, err, redis.Nil)

	// set and get
	rds.Set(ctx, key, true, 30*time.Minute)
	val, err := rds.Get(ctx, key).Bool()
	assert.NotEqual(t, err, redis.Nil)
	assert.Equal(t, val, true)

	// set again
	rds.Set(ctx, key, false, 30*time.Minute)
	val, err = rds.Get(ctx, key).Bool()
	assert.NotEqual(t, err, redis.Nil)
	assert.Equal(t, val, false)
}

func TestHashSet(t *testing.T) {
	hand := New(getRds())
	rds := hand.GetClient()
	key := RdKey(`hashexp:%d`, time.Now().Unix())

	// err := rds.HMSet(ctx, "map",
	// 	"name", "hello",
	// 	"count", 123,
	// 	"correct", true).Err()
	// if err != nil {
	// 	panic(err)
	// }
	//
	// // Get the map. The same approach works for HmGet().
	// res := rds.HGetAll(ctx, "map")
	// if res.Err() != nil {
	// 	panic(err)
	// }
	//
	// type data struct {
	// 	Name    string `redis:"name"`
	// 	Count   int    `redis:"count"`
	// 	Correct bool   `redis:"correct"`
	// }
	//
	// // Scan the results into the struct.
	// var d data
	// if err := res.Scan(&d); err != nil {
	// 	panic(err)
	// }
	//
	// t.Log(d)

	// hgetall
	var us UserList
	ma := rds.HGetAll(ctx, key).Val()
	for k, v := range ma {
		t.Logf("k=>%v, v=>%+v", k, v)
		uv := &User{}
		err := json.Unmarshal([]byte(v), uv)
		assert.Nil(t, err)
		us = append(us, uv)
	}

	// hsetall
	// mus := []*User{
	// 	{78, "clark"},
	// 	{32, "terry"},
	// 	{90, "raft"},
	// }
	// u := &User{77, "gght"}
	// set := rds.HMSet(ctx, key, u)
	// assert.Nil(t, set.Err())

	// hgetall
	var usm UserList
	ma = rds.HGetAll(ctx, key).Val()
	for k, v := range ma {
		t.Logf("k=>%v, v=>%+v", k, v)
		uv := &User{}
		err := json.Unmarshal([]byte(v), uv)
		assert.Nil(t, err)
		usm = append(usm, uv)
	}
	t.Logf("usm:+%+v", usm)
}

func TestSetBytes(t *testing.T) {
	// write to redis
	u := &User{1, "clark"}
	data, err := json.Marshal(u)
	assert.Nil(t, err)

	// read from redis
	gu := &User{}
	err = json.Unmarshal(data, gu)
	if err != nil {
		t.Errorf("got err, %s1", err)
	}
	t.Logf("user %+v", gu)

	// set
	rds := getRds()
	ctx := context.Background()
	key := "user:1"
	rds.Set(ctx, key, string(data), 100*time.Second)

	// get
	key1, key2 := key, "user:no"
	s1 := rds.Get(ctx, key1).String()	// exist
	t.Logf("get(key1)=%+v, empty=%t", s1, s1=="")
	s2 := rds.Get(ctx, key2).String() // not exist
	t.Logf("get(key2)=%+v, empty=%t", s2, s2=="redis:nil")
	exists := rds.Exists(ctx, key1, key2)
	t.Logf("exist %+v", exists)

	// cmd
	sc1 := rds.Get(ctx, key1)
	sc2 := rds.Get(ctx, key2)

	t.Logf("sc1, Err()=%v, Val=%s, String=%s", sc1.Err(), sc1.Val(), sc1.String())
	t.Logf("sc2, Err()=%v, Val=%s, String=%s, sc2.Err() == redis.Nil got %t", sc2.Err(), sc2.Val(), sc2.String(), sc2.Err() == redis.Nil)

	// del
	rds.Set(ctx, "k","2", -1)
	dkey1 := rds.Del(ctx, key1, "k")
	dkey2 := rds.Del(ctx, key2, "k")
	t.Logf("del key1, Err()=%v, Val=%d", dkey1.Err(), dkey1.Val())
	t.Logf("del key2, Err()=%v, Val=%d", dkey2.Err(), dkey2.Val())


	return
}

type User struct {
	ID   uint64 `redis:"id"`
	Name string `redis:"name"`
}

type UserList []*User

func (u *UserList) MarshalBinary() (data []byte, err error) {
	return json.Marshal(u)
}
