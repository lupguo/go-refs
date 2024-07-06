package errgroupx

import (
	"math/rand"
	"testing"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func fn(t *testing.T, id int, duration time.Duration) error {
	t.Logf("exec fn(%d)\n", id)
	randNum := rand.Int31n(100)

	if duration == 0 {
		time.Sleep(time.Duration(randNum) * time.Millisecond)
	} else {
		time.Sleep(duration)
	}

	if randNum > 70 {
		return errors.Errorf("func(%d) over 60, randNum=%d", id, randNum)
	}
	return nil
}

func TestErrGroupx(t *testing.T) {
	egroup := errgroup.Group{} // egroup.Go报错，

	// g1
	egroup.Go(func() error {
		return fn(t, 1, 0)
	})
	// g2
	egroup.Go(func() error {
		return fn(t, 2, 0)
	})
	// g3
	egroup.Go(func() error {
		return fn(t, 3, time.Second)
	})

	// errgroup.Wait()会等待之前加入的Go协程返回完成，同时仅返回第一个有错误的消息
	if err := egroup.Wait(); err != nil {
		t.Logf("egroup wait got err: %v\n", err)
	}

	t.Log("test finished!")
}
