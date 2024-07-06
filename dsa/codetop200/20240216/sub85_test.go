package _0240216

import (
	"testing"
)

// 有效括号生成
// n=2 => (())、()()
// 思路: 回溯思想
//  1. 有两种选择 [(,)],
//  2. 满足条件 :
//     a. ( 总数没有超过 n，超过回退选择
//     b. ( 总数没有超过 n, 当前path得到剩余(的数量>0，才可以选择 [(,)]，否则只能选择[(]
func generateParenthesis(n int) []string {
	var leftUsed int // 使用了多少个左括号，
	var path string  // 路径解
	var res []string // 可行解
	parenthesisBacktrack(n, leftUsed, path, &res)
	return res
}

func parenthesisBacktrack(n int, leftUsed int, path string, res *[]string) {
	// 找到一个可行解
	if leftUsed == n && len(path) == 2*n {
		*res = append(*res, path)
		return
	}

	for _, choice := range []string{"(", ")"} {
		// 排除异常情况
		if leftUsed > n || !parenthesisMatch(path, choice) {
			continue
		}

		// 选择left choice
		if choice == "(" {
			leftUsed++
		}
		path += choice

		// 递归到下层决策树
		parenthesisBacktrack(n, leftUsed, path, res)

		// 回退
		if choice == "(" {
			leftUsed--
		}
		path = path[:len(path)-1]
	}
}

// 检测path加上choice括号后是否满足合格括号规则
func parenthesisMatch(path string, choice string) bool {
	var noMatch int
	for i := 0; i < len(path); i++ {
		if path[i] == '(' {
			noMatch++
		} else {
			noMatch--
		}
	}

	if noMatch == 0 && choice == ")" {
		return false
	}

	return true
}

func TestGenerateParenthesis(t *testing.T) {
	t.Logf("generateParenthesis(2)=%v", generateParenthesis(2))
	t.Logf("generateParenthesis(3)=%v", generateParenthesis(3))
}
