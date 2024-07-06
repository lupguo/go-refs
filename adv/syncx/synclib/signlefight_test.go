package synclib

import (
	"sync"
	"testing"
	"time"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
)

func TestSightFight(t *testing.T) {
	sfg := singleflight.Group{}

	// key提前确认
	fakeRpc := func() (string, error) {
		time.Sleep(1 * time.Second)
		return "ok", nil
	}

	type result struct {
		val    interface{}
		err    error
		shared bool
		id     int
	}

	// 并发call
	sfgKey := "createGroup"
	smap := sync.Map{}
	eg := errgroup.Group{}
	for i := 0; i < 10; i++ {
		id := i // 注意闭包问题
		eg.Go(func() error {
			// 模拟在不同协程并发rpc call
			val, err, shared := sfg.Do(sfgKey, func() (interface{}, error) {
				return fakeRpc()
			})

			// mutex insert
			smap.Store(id, result{
				id:     id,
				val:    val,
				err:    err,
				shared: shared,
			})

			return err
		})
	}

	// 并发中有任何返回的错误，则终止并发返回
	if err := eg.Wait(); err != nil {
		t.Logf("eg got err: %s", err)
		return
	}

	// 读取
	smap.Range(func(key, value any) bool {
		t.Logf("sfgKey=>%d, val=>%+v\n", key, value)
		return true
	})

}
