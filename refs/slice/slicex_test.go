package slice

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 模拟字符串栈出栈操作，不过注意s实际是切片的值拷贝，这里要返回新的栈地址
func TopStrStack(s []string) (top string, stk []string) {
	n := len(s)
	if n > 0 {
		top = s[n-1]
		n--
		s = s[:n]
	}
	return top, s
}

// 更新字符串栈中原始值，重复一次
func UpdateStrStack(s []string) {
	for i, v := range s {
		s[i] = strings.Repeat(v, 2)
	}
}

// 通过函数向底层Slice追加元素，观察切片长度和容量变化
func AppendStrStack(stk []string, n int) (int, int, []string) {
	rawLen := len(stk)
	rawCap := cap(stk)
	for i := 0; i < n; i++ {
		stk = append(stk, strconv.Itoa(i))
	}
	// cap 扩容后，更改原始的初始值
	if cap(stk) != rawCap {
		for i := 0; i < rawLen; i++ {
			stk[i] = fmt.Sprintf("By inner cap expand upd:%v", stk[i])
		}
	}
	return len(stk), cap(stk), stk
}

func TestAppendStrStack(t *testing.T) {
	tests := []struct {
		name string
		stk  []string
	}{
		{"t0", []string{}},
		{"t1", []string{"a"}},
		{"t2", []string{"a", "b"}},
		{"t3", []string{"a", "b", "c"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lenN, capN, stk := AppendStrStack(tt.stk, 10)
			t.Logf("rawLenN=%v, rawCapN=%v, raw tt.stk=%v", len(tt.stk), cap(tt.stk), tt.stk)
			t.Logf("lenN=%v, capN=%v, stk=%v ", lenN, capN, stk)
		})
	}
}

func TestTopStringStack(t *testing.T) {
	tests := []struct {
		name string
		stk  []string
		want string
	}{
		{"t1", []string{"a", "b", "c"}, "c"},
		{"t2", []string{"a", "b"}, "b"},
		{"t3", []string{"a"}, "a"},
		{"t4", []string{}, ""},
		{"t5", nil, ""},
	}
	for _, tt := range tests {
		inner := tt.stk
		t.Run(tt.name, func(t *testing.T) {
			got, newStk := TopStrStack(inner)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
			t.Logf("tt.stk=%v, stk=%v", tt.stk, newStk)
		})
	}
}

func TestUpdateStrStack(t *testing.T) {
	tests := []struct {
		name string
		stk  []string
	}{
		{"t1", []string{"a", "b", "c"}},
		{"t2", []string{"a", "b"}},
		{"t3", []string{"a"}},
		{"t4", []string{}},
		{"t5", nil},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			UpdateStrStack(tt.stk)
			t.Logf("tt.stk=%v", tt.stk)
		})
	}
}

func TestSlice(t *testing.T) {
	nums := []int{1, 2, 3, 4, 5} // len=5
	t.Log(nums[0:len(nums)])
	t.Log(nums[0:])
	// t.Log(nums[0:6]) // slice bounds out of range [:6] with capacity 5
	t.Log(nums[0 : len(nums)-1])
	t.Log(nums[:len(nums)-1])
	t.Log(nums[:])
	t.Log(nums[:1])
	t.Log(nums[:0]) // 从idx=0开始，增加0个长度，左开右闭区间
	t.Log(nums[0:0])
	t.Log(nums[1:])
	t.Log(nums[1:2]) // 从idx=1开始，到idx=2结束(不包含），左开右闭区间[1,2)
	t.Log(nums[5:])  // 从位置5开始，返回nil
	t.Log(nums[5:5]) // 从位置5开始，返回nil
	t.Log(nums[6:])  // 区间溢出，slice bounds out of range [6:5] [recovered] slice bounds out of range [6:5]
}
