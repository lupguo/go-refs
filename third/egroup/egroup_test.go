package egroup

import (
	"testing"

	"golang.org/x/sync/errgroup"
)

func TestEgroup(t *testing.T) {
	fn := func(id int) {
		// 在fn内部recover
		defer func() {
			if r := recover(); r != nil {
				t.Logf("recover in fn: %v", r)
			}
		}()

		t.Logf("id=%v", id)
		if id == 3 {
			panic("id[3] panic")
		}
	}

	// errgroup 部分
	egroup := errgroup.Group{}
	for i := 0; i < 5; i++ {
		i := i // 闭包问题
		egroup.Go(func() error {
			fn(i)
			return nil
		})
	}

	if err := egroup.Wait(); err != nil {
		t.Errorf("egroup got err: %s", err)
	}

	t.Logf("done")
}
