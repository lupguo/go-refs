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

func TestEgroup2(t *testing.T) {
	// errgroup 部分
	egroup := errgroup.Group{}
	egroup.SetLimit(3)
	for i := 0; i < 10; i++ {
		i := i // 特别注意，闭包问题
		egroup.Go(func() error {
			t.Logf("#%d", i)
			return nil
		})
	}

	if err := egroup.Wait(); err != nil {
		t.Errorf("egroup got err: %s", err)
	}

	t.Logf("done")
}

func TestEgroup3(t *testing.T) {
	a, b, c := "a", "b", "c"
	list := []*string{&a, &b, &c}
	// errgroup 部分
	egroup := errgroup.Group{}
	for k, v := range list {
		k, v := k, v // 特别注意，闭包问题
		egroup.Go(func() error {
			t.Logf("#%d, v=>%v, val v=>%v", k, v, *v)
			return nil
		})
	}

	if err := egroup.Wait(); err != nil {
		t.Errorf("egroup got err: %s", err)
	}

	t.Logf("done")
}

func BenchmarkGo(b *testing.B) {
	fn := func() {}
	g := &errgroup.Group{}
	b.ResetTimer()
	b.ReportAllocs()
	for i := 0; i < b.N; i++ {
		g.Go(func() error { fn(); return nil })
	}
	g.Wait()
}
