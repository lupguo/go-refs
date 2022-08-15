package defert

import (
	"fmt"
	"testing"
	"time"
)

// defer 符合FILO栈的运行方式，defer函数会压栈

func TestDefer(t *testing.T) {
	a := 1

	// 压入
	defer func() {
		fmt.Printf("a1=%d", a) // a1=
	}()

	// 压入
	defer func(a int) {
		fmt.Printf("a2=%d", a) // a2=1 赋值传入了，a不受最后的a++影响
	}(a)

	// 压入
	defer func(a *int) {
		*a++
		fmt.Printf("a3=%d", *a) // a3=
	}(&a)

	if true {
		// 压入
		defer fmt.Println(1) //
	}

	fmt.Println("func p1") // 打印 func p1

	defer fmt.Println(2) // 虽然是传入，fmt.Println也是出栈时候执行，不必使用func(){}再包一层

	defer fmt.Println(3)

	a++
}

func f(t *testing.T) {

}

func g(t *testing.T, n int) {
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		t.Logf("g recover: %v", r)
	// 	}
	// }()

	t.Logf("g[%d] exec..", n)
	if n == 2 {
		g2()
	}
	t.Logf("g[%d] success!!", n)
}

func g2() {
	panic("g2 panic")
}

func TestGoDeferPanic(t *testing.T) {
	// 注意，如果是Goroutine没有recover，则会直接panic掉
	// 下面的代码只能recover当前goroutine的的崩溃
	// defer func() {
	// 	if r := recover(); r != nil {
	// 		t.Logf("recover: %v", r)
	// 	}
	// }()
	// 带协程执行，内部有panic，但在协程内，panic发起go并发的协程没有办法处理做恢复，必须在go协程内部恢复
	for i := 0; i < 3; i++ {
		go g(t, i)
	}
	time.Sleep(1 * time.Second)
	t.Logf("main done!")
}
