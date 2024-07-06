package _0240216

import (
	"testing"
)

// N皇后问题
// https://leetcode.cn/problems/n-queens/
// 思路: 回溯算法
//  1. n个皇后，可行解每行至少一个，否则存在一行2个，则必定冲突
//  2. path路径为每行选择，res为结果，选择排放位置..Q.
func solveNQueens(n int) [][]string {
	var res [][]string
	var path []string

	queenBacktrack(n, 0, path, &res)
	return res
}

func queenBacktrack(n, row int, path []string, res *[][]string) {
	if len(path) == n {
		*res = append(*res, path)
		return
	}

	for col := 0; col < n; col++ { // 行
		// 检测
		if conflictQueenPut(path, row, col) {
			continue
		}

		// 行选择
		queenStr := rowQueenStr(col, n)
		path = append(path, queenStr)

		// 递归下层决策树
		row++
		queenBacktrack(n, row, path, res)

		// 行回退
		row--
		path = removePaths(path, queenStr)
	}
}

func removePaths(path []string, str string) []string {
	var res []string
	for _, s := range path {
		if s != str {
			res = append(res, s)
		}
	}
	return res
}

func conflictQueenPut(path []string, row int, col int) bool {
	for x, qstr := range path {
		y := rowQueenCol(qstr)

		// 行、列冲突
		if x == row || y == col {
			return true
		}

		// 斜线冲突
		if x+y == row+col || x-y == row-col {
			return true
		}
	}

	return false
}

func rowQueenCol(qstr string) int {
	for i, v := range qstr {
		if v == 'Q' {
			return i
		}
	}
	return 0
}

func rowQueenStr(col int, n int) string {
	var s string
	for i := 0; i < n; i++ {
		if i == col {
			s += "Q"
		} else {
			s += "."
		}
	}
	return s
}

func TestSolveNQueens(t *testing.T) {
	t.Logf("4x4 Queue: %v", solveNQueens(4))
	t.Logf("5x5 Queue: %v", solveNQueens(5))
	t.Logf("8x8 Queue: %v", solveNQueens(8))
}
