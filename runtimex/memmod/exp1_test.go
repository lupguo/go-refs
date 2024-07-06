package memmod

import (
	"testing"
	"time"
)

// 可能输出
//  exp1_test.go:21: memmod catch a bug: a=0, b=2
//  exp1_test.go:21: memmod catch a bug: a=1, b=2
func TestExp1(t *testing.T) {
	for i := 0; i < 1e6; i++ {
		var a, b int

		f := func() {
			a = 1  // w之前的写操作
			b = 10 // 写操作w
		}

		g := func() {
			for b != 10 { // 确保b==10
			}

			if a != 1 && a != 2 {
				t.Logf("memmod catch a bug: a=%v, b=%v", a, b) // a可能为0，1
			}
		}

		go f() // g1
		g()    // g2

	}
}

func TestExp1A(t *testing.T) {
	v := 0
	f := func() {
		println(v)
	}
	for i := 0; i < 5; i++ {
		v = i
		go f()
	}
	time.Sleep(time.Second)
}
