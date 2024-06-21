package leetcode

//
// func TestLCOF001(t *testing.T) {
// 	var
// 	tests := []struct {
// 		a int
// 		b int
// 	}{{}}
// }

func divide(a int, b int) int {
	// 异常处理
	if a == 0 || b == 0 {
		return 0
	}

	// 符号处理
	sign := 0
	if a < 0 {
		a = -a
		if b < 0 {
			b = -b
		} else {
			sign = 1
		}
	} else if b < 0 {
		b = -b
		sign = 1
	}

	// 循环除法设定
	var v int
	for a-b >= 0 {
		v++
		a = a - b
	}

	// 边界检测
	min, max := -1<<31, 1<<31-1
	if v <= min || v >= max {
		return max
	}

	// 符号处理
	if sign == 1 {
		return -v
	}
	return v
}
