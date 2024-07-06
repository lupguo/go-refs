package _0231124

import (
	"fmt"
	"strconv"
	"strings"
	"testing"
)

// 给定一个经过编码的字符串，返回它解码后的字符串。
//
// 编码规则为: k[encoded_string]，表示其中方括号内部的 encoded_string 正好重复 k 次。注意 k 保证为正整数。
//
// 你可以认为输入字符串总是有效的；输入字符串中没有额外的空格，且输入的方括号总是符合格式要求的。
//
// 此外，你可以认为原始数据不包含数字，所有的数字只表示重复的次数 k ，例如不会出现像 3a 或 2[4] 的输入。
// 示例 1：
//
// 输入：s = "3[a]2[bc]"
// 输出："aaabcbc"
// 示例 2：
//
// 输入：s = "3[a2[c]]"
// 输出："accaccacc"
// 示例 3：
//
// 输入：s = "2[abc]3[cd]ef"
// 输出："abcabccdcdcdef"
// 示例 4：
//
// 输入：s = "abc3[cd]xyz"
// 输出："abccdcdcdxyz"
func decodeString(s string) string {

	return ""
}

// 思路:
//
//	从左向右迭代:
//	1. 遇到数字记录到num(注意连续多位数字情况)，待后续存储到 numStack 用于倍数处理
//	2. 遇到字符存入 str(注意连续存在多位情况)，待后续存入到 strStack(可能为""字符，也需要入栈)
//	3. 遇到[符号，表示[字符之前累计的数字\字符串可以入栈到 numStack，strStack
//	4. 遇到]符号: 当前的str为[]重复内容, 从strStack出栈一个用于连接前缀 + Strings.Repeat(str, num)，将拼接的内容存入str中
//	 str = strStack.Pop() + Strings.Repeat(str, numStack().Pop())
//	 str 就是要求解的值
func DecodeStringMethod02(s string) string {
	var strStack []string
	var numStack []int

	var times int
	var str string
	for _, c := range s {
		switch {
		case c >= '0' && c <= '9': // 字符转10进制数字
			num, _ := strconv.Atoi(string(c))
			times = times*10 + num
		case c >= 'a' && c <= 'z' || c >= 'A' && c <= 'Z': // 组装字符串，可能是[之前的前缀字符串，也可能是[内的innerStr字符串
			str += string(c)
		case c == '[':
			strStack = append(strStack, str) // 将 [ 的前缀字符串暂存入栈
			str = ""
			numStack = append(numStack, times) // 将[ 之前的倍数暂存入栈
			times = 0
		case c == ']': // 出栈，出栈时候，str 定义就是待[str]内重复字符串内容
			innerStr := str

			// 重复次数出栈
			repTimes, newNumStack := Pop(numStack)
			numStack = newNumStack

			// 前缀字符串
			preStr, newStrStack := Pop(strStack)
			strStack = newStrStack

			// 拼接成内部串，存入str存入结果集合
			str = preStr + strings.Repeat(innerStr, repTimes)
		}
	}

	return str
}

func Pop[S []E, E any](stk S) (E, S) {
	topElem := stk[len(stk)-1]
	stk = stk[:len(stk)-1]
	return topElem, stk
}

func TestDecodeStringMethod01(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"t1", "3[a]2[bc]", "aaabcbc"},
		{"t2", "3[a2[c]]", "accaccacc"},
		{"t3", "2[abc]3[cd]ef", "abcabccdcdcdef"},
		{"t4", "abc3[cd]xyz", "abccdcdcdxyz"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DecodeStringMethod02(tt.s); got != tt.want {
				t.Errorf("got %s, but want %s", got, tt.want)
			}
		})
	}
}

func TestPop(t *testing.T) {
	type TElem interface {
		string | int
	}

	tests := []struct {
		name     string
		stk      []any
		wantElem any
	}{
		{"t1", []any{1, 2, 3}, 3},
		{"t1", []any{"a", "b", "c"}, "c"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			t.Logf("pointer(&tt.stk)=%p", &tt.stk)
			got, nstk := Pop(tt.stk)
			if got != tt.wantElem {
				t.Errorf("got %v, but want %v", got, tt.wantElem)
			}
			t.Logf("tt.stk=%v, nstk=%v", tt.stk, nstk)
		})
	}

}

// 思路: 利用栈的特性每当遇到]寻找对应的[，其中为待重复字符串，[之前的数字则为k次
// eg. 3[a2[c]]
//  1. 如何存储数据? -> (3, [, a,  2, [, c, ] , ]
//     a) stack: (bef="",mul=3, sub=?), ( bef=a, mul=2, sub=[c])
//     b) stack: 3, [, a, 2, [, c, ]
//     build(k,s)
func DecodeStringMethod01(s string) string {
	var stk []string
	for i := 0; i < len(s); i++ {
		r := s[i]
		switch {
		case r >= '0' && r <= '9': // 获取重复数字
			numStr, j := getRepTimes(s, i) // num
			stk = append(stk, numStr)
			i = j
		case r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z': // 前缀字符串
			subStr, j := getSubStr(s, i) // string
			stk = append(stk, subStr)
			i = j
		case r == '[': // 左括号
			stk = append(stk, string(r))
		case r == ']': // 出栈
			stk = stkPopAndPush(stk)
		}
	}

	return stk[0]
}

// 遇到右括号]，做出栈操作，出栈到下个左括号[或者空栈截止
func stkPopAndPush(stk []string) (newStk []string) {
	var foundBracket int
	n := len(stk)
	var pops []string
	top := n - 1
	for top >= 0 {
		if foundBracket > 0 {
			break
		}

		if stk[top] == "[" {
			foundBracket++
			continue
		}
		pops = append(pops, stk[top])
		top--
	}

	if top == -1 {
		top = 0
	}

	// 更新stk栈
	var innerStr string
	switch len(pops) {
	case 1:
		innerStr = pops[0]
	case 2:
		mul, _ := strconv.Atoi(pops[1])
		innerStr = strings.Repeat(pops[0], mul)
	case 3:
		mul, _ := strconv.Atoi(pops[1])
		innerStr = fmt.Sprintf("%s%s", pops[2], strings.Repeat(pops[0], mul))
	}

	stk = append(stk[:top], innerStr)
	return stk
}

// 从字符串s中寻找从start开始的子串
func getSubStr(s string, start int) (subStr string, end int) {
	i := start
	for ; i < len(s); i++ {
		r := s[i]
		if r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' {
			subStr += string(r)
			continue
		}
		break
	}
	return subStr, i - 1
}

// 从字符串中s寻找第start位置，组装成数字返回
func getRepTimes(s string, start int) (num string, end int) {
	var numStr string
	i := start
	for ; i < len(s); i++ {
		if s[i] >= '0' && s[i] <= '9' {
			numStr += string(s[i])
			continue
		}
		break
	}

	return numStr, i - 1
}
