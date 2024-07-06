package syncxmap

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/pkg/errors"
	"golang.org/x/sync/errgroup"
)

func TestEGPGo(t *testing.T) {
	egp := errgroup.Group{}
	// egp.SetLimit(10)

	var sum1, sum2, iterSum int
	for i := 0; i < 100; i++ {
		j := i
		iterSum += i
		egp.Go(func() error {
			sum1 += j
			if j == 50 {
				return errors.New("egp go func j=50, mock err")
			}
			sum2 += j
			return nil
		})
	}

	// 等待所有的egp.Go()协程返回,
	err := egp.Wait()
	if err != nil {
		// 按理sum2 和iterSum接近，因为egp.Go协程是并发的，其中一个返回(j=50)，并不影响剩余的egp.Go协程的sum2+=j操作
		// err为首个egp.Go返回的错误
		t.Errorf("egp got err: %v, sum1=%v, sum2=%v", err, sum1, sum2)
		return
	}

	t.Logf("sum=%v, iter sum=%v", sum1, iterSum)
}

func TestEGPTryGo(t *testing.T) {
	egp := errgroup.Group{}
	// egp.SetLimit(10)

	var sum, iterSum int
	for i := 0; i < 100; i++ {
		j := i
		iterSum += i
		egp.TryGo(func() error {
			sum += j
			if j == 50 {
				return errors.New("egp go func j=50, mock err")
			}
			return nil
		})
	}

	err := egp.Wait()
	if err != nil {
		t.Errorf("egp got err: %v", err)
		return
	}

	t.Logf("sum=%v, iter sum=%v", sum, iterSum)
}

func egpCtx() {
	g, ctx := errgroup.WithContext(context.Background())

	for i := 0; i < 10; i++ {
		index := i // 避免闭包引用相同变量
		g.Go(func() error {
			return doEgpCtxFn(ctx, index)
		})
	}

	if err := g.Wait(); err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("All tasks completed successfully")
	}
}

func doEgpCtxFn(ctx context.Context, index int) error {
	fmt.Printf("Task %d started\n", index)

	select {
	case <-ctx.Done():
		return ctx.Err()
	default:
		// 模拟耗时操作
		time.Sleep(time.Second * time.Duration(index))

		if index == 5 { // 模拟一个任务失败的情况
			return fmt.Errorf("Task %d failed", index)
		}

		fmt.Printf("Task %d completed\n", index)
		return nil
	}
}

func BenchmarkEGPSum(b *testing.B) {
	for i := 0; i < b.N; i++ {

	}
}
