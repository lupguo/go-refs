package _0231011

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/reverse-string/description/
// 编写一个函数，其作用是将输入的字符串反转过来。输入字符串以字符数组 s 的形式给出。
//
// 不要给另外的数组分配额外的空间，你必须原地修改输入数组、使用 O(1) 的额外空间解决这一问题。
func reverseString(s []byte) {
	// 两边i,j夹逼, i++, j--
	i := 0
	j := len(s) - 1

	// 当i < j 时候进行两个元素交换
	for i < j {
		s[i], s[j] = s[j], s[i]
		i++
		j--
	}
}

func TestReverseString(t *testing.T) {
	tests := []struct {
		name string
		s    []byte
		want []byte
	}{
		{"t1", []byte{'a'}, []byte{'a'}},
		{"t2", []byte{'a', 'b'}, []byte{'b', 'a'}},
		{"t3", []byte{'a', 'b', 'c'}, []byte{'c', 'b', 'a'}},
		{"t4", []byte{'a', 'b', 'c', 'd'}, []byte{'d', 'c', 'b', 'a'}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			reverseString(tt.s)
			if !reflect.DeepEqual(tt.s, tt.want) {
				t.Errorf("got %v, but want %v", tt.s, tt.want)
			}
		})
	}

}
