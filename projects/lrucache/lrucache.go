package lrucache

import (
	"sync"
	"time"

	"github.com/pkg/errors"
)

type Cache interface {
	Set(key string, v interface{}, ttl time.Duration) error

	Get(key string) (v interface{}, err error)

	Del(key string)
}

type data struct {
	val        interface{}
	expireTime time.Time
	weight     uint64
}

type activeLink struct {
	next *data
}

type LRUCache struct {
	lock       sync.RWMutex
	store      map[string]*data
	activeKeys *activeLink
}

func (c *LRUCache) Set(key string, v interface{}, ttl time.Duration) error {
	c.store[key] = &data{
		val:        v,
		expireTime: time.Now().Add(ttl),
	}
	return nil
}

func (c *LRUCache) Get(key string) (v interface{}, err error) {
	data, ok := c.store[key]
	if !ok {
		return nil, errors.New("key not exist")
	} else if data.expireTime.Sub(time.Now()) <= 0 {
		return nil, errors.New("key is out of date")
	}

	// 访问权重增加
	data.weight++

	return data, nil
}

func (c *LRUCache) Del(key string) {
	// TODO implement me
	panic("implement me")
}
