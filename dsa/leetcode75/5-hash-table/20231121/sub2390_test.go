package _0231121

import (
	"testing"
)

// 给你一个包含若干星号 * 的字符串 s 。
//
// 在一步操作中，你可以：
//
// 选中 s 中的一个星号。
// 移除星号 左侧 最近的那个 非星号 字符，并移除该星号自身。
// 返回移除 所有 星号之后的字符串。
//
// 注意：
//
// 生成的输入保证总是可以执行题面中描述的操作。
// 可以证明结果字符串是唯一的。
//
// 输入：s = "leet**cod*e"
// 输出："lecoe"
// 解释：从左到右执行移除操作：
// - 距离第 1 个星号最近的字符是 "leet**cod*e" 中的 't' ，s 变为 "lee*cod*e" 。
// - 距离第 2 个星号最近的字符是 "lee*cod*e" 中的 'e' ，s 变为 "lecod*e" 。
// - 距离第 3 个星号最近的字符是 "lecod*e" 中的 'd' ，s 变为 "lecoe" 。
//
// 示例 2：
//
// 输入：s = "erase*****"
// 输出：""
// 解释：整个字符串都会被移除，所以返回空字符串
func removeStars(s string) string {

	return ""
}

// 思路: 一看就有栈的味道了
// 1. 从0开始迭代s长度，如果遇到*，pop出一个元素
// 2. 返回栈内剩余字符
func RemoveStarsMethod01(s string) string {
	// 1. Go中栈实现 slice? 入栈 s.append()操作,s.len+1，出栈s.len-1
	// 2. 最后返回ss[0:s.len]
	var ss []rune
	var n int
	for _, r := range s {
		if r == '*' {
			if n >= 1 {
				ss = ss[:n-1]
				n--
			}
		} else {
			ss = append(ss, r)
			n++
		}
	}

	return string(ss[:n])
}

func TestRemoveStarsMethod01(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"t1", "leet**cod*e", "lecoe"},
		{"t2", "erase*****", ""},
		{"t3", "***", ""},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveStarsMethod01(tt.s); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
