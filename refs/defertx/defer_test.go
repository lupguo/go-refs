package defertx

import (
	"testing"
	"time"
)

// defer 符合FILO栈的运行方式，defer函数会压栈
func TestDefer(t *testing.T) {
	a := 1

	// 压入，闭包
	defer func() {
		t.Logf("a1=%d", a) // a1=(a的值受后续a的值变化影响)
	}()

	// 压入
	defer func(a int) {
		t.Logf("a2=%d", a) // a2=1 赋值传入了，a不受最后的a++影响
	}(a)

	// 压入
	defer func(a *int) {
		*a = *a + 100
		t.Logf("a3=%d", *a) // 传值，函数内部对a的值改变，会影响其他函数对a值的读取
	}(&a)

	if true {
		// 压入
		defer t.Log(1) //
	}

	t.Log("func p1") // 打印 func p1

	defer t.Log(2) // 虽然是传入，t.Log也是出栈时候执行，不必使用func(){}再包一层
	defer t.Log(3)

	defer t.Logf("a=%d", a) // a=1

	a++
}

// a=0
func TestFn1(t *testing.T) {
	var a = 0
	defer t.Logf("not a=%d", a) //  a=0 -> 立马进行值拷贝
	a++
	defer t.Logf("fn a=%d", a) //   a=1 -> 立马进行值拷贝
	a++
}

// a=1
func TestFn2(t *testing.T) {
	var a = 0
	defer func() {
		t.Logf("a=%d", a) // a=2 -> 最后执行func()时候进行值拷贝
	}()
	a++
	defer func() {
		t.Logf("a=%d", a) // a=2 -> 最后执行func()时候进行值拷贝
	}()
	a++
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
