package memmod

import (
	"testing"
)

type M struct {
	msg string
}

// 输出
//	example3_test.go:30: memmod catch bug, &x=0x14000180150, &y=0x14000180150, x=&{msg:hello}, y=&{msg:hello}
//  example3_test.go:30: memmod catch bug, &x=0x14000180220, &y=0x14000180220, x=&{msg:hello}, y=&{msg:hello}
//  example3_test.go:30: memmod catch bug, &x=0x14000180780, &y=0x14000180780, x=&{msg:hello}, y=&{msg:hello}
//  example3_test.go:30: memmod catch bug, &x=0x14000180880, &y=0x14000180880, x=&{msg:hello}, y=&{msg:hello}
func TestExample3(t *testing.T) {

	// g1协程
	for i := 0; i < 1e3; i++ {
		var x, y *M

		// 起一个协程g2
		go func() {
			x = new(M)
			x.msg = "hello"
			y = x // x赋值给y
		}()

		// g1 检测y为空则，g1不断轮询等待
		for y == nil {
		}

		// g1 检测到y不为空，但y.msg仍然有可能不为hello
		// 即应该注意: 不同协程不保证按代码顺序执行
		if y.msg != "hello" {
			t.Errorf("memmod catch bug, &x=%p, &y=%p, x=%+v, y=%+v", x, y, x, y)
		}
	}

	t.Log("done")
}
