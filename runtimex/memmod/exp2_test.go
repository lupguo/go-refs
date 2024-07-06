package memmod

import (
	"testing"
)

func TestExp2(t *testing.T) {
	for i := 0; i < 1e9; i++ {
		var a string

		f := func() {
			// print(a)
			if a != "hello" {
				t.Errorf("memod got err, a=%v", a)
			}
		}

		q := func() {
			a = "hello"
			go f() // 必定会输出hello
		}

		q()
	}

}
