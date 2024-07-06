package syncxmap

import (
	"errors"
	"fmt"
	"testing"

	"golang.org/x/sync/errgroup"
	"golang.org/x/sync/singleflight"
)

func TestSingleFight(t *testing.T) {
	sf := singleflight.Group{}
	eg := errgroup.Group{}
	eg.SetLimit(1000)

	count := 0
	for i := 0; i < 1e4; i++ {
		key := fmt.Sprintf("sf:%v", i)
		eg.Go(func() error {
			// 使用singleflight
			ret, err, shard := sf.Do(key, func() (any, error) {
				// exec once
				count++
				return count, nil
			})
			if err != nil {
				t.Errorf("got err: %v", err)
				return errors.Join(errors.New("eg Go got err"), err)
			}
			t.Logf("key=%v, ret=%v, b=%v", key, ret, shard)

			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		t.Errorf("got err: %v", err)
	}

}

func TestSingleFight1(t *testing.T) {
	// sf := singleflight.Group{}
	eg := errgroup.Group{}
	eg.SetLimit(1000)

	count := 0
	for i := 0; i < 1e4; i++ {
		key := fmt.Sprintf("sf:%v", i)
		eg.Go(func() error {
			count++
			t.Logf("key=%v, b=%v", key, count)
			return nil
		})
	}

	if err := eg.Wait(); err != nil {
		t.Errorf("got err: %v", err)
	}

}
