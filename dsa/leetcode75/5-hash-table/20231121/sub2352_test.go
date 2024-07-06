package _0231121

import (
	"fmt"
	"testing"
)

// 给你一个下标从 0 开始、大小为 n x n 的整数矩阵 grid ，返回满足 Ri 行和 Cj 列相等的行列对 (Ri, Cj) 的数目。
//
// 如果行和列以相同的顺序包含相同的元素（即相等的数组），则认为二者是相等的。
//
// https://leetcode.cn/problems/equal-row-and-column-pairs/description/?envType=study-plan-v2&envId=leetcode-75
func equalPairs(grid [][]int) int {

	return 0
}

// 思路: 实际比较行\列是否一致(通过hash计数）-> 相等类问题优先考虑map
// 1. 初始化行、列的hash map, mapR、mapC -> map[string]int
// 2. 迭代mapR，基于检测是否在mapC中存在，若存在，取max(R,C)中的计数最大值累加(res)
//
// AC过程几个问题:
// 1. 没有考虑Map的key使用数字拼接问题(需要通过:分割)
// 2. 二维数组的行列迭代，可以通过两层循环分别使用num[i][j]和mum[j][i] 识别
// 3. 算法计数问题，两个map是通过v1*v2得到，而不是max(v1,v2) -> 组合问题->可以重复匹配
func EqualPairsMethod01(grid [][]int) int {
	// 1. 初始化行、列map
	mapR, mapC := make(map[string]int), make(map[string]int)
	// grid[0][0] grid[1][0] ...grid[n-1][0]
	n := len(grid)
	for i := 0; i < n; i++ {
		var colStr, rowStr string
		for j := 0; j < n; j++ {
			rowStr = fmt.Sprintf("%s:%d", rowStr, grid[i][j])
			colStr = fmt.Sprintf("%s:%d", colStr, grid[j][i])
		}
		mapR[rowStr]++
		mapC[colStr]++
	}

	// 2. 比较行列map值是否一致
	var cnt int
	for str, cntR := range mapR {
		if cntC, ok := mapC[str]; ok {
			cnt += cntC * cntR
		}
	}

	return cnt
}

func TestEqualPairsMethod01(t *testing.T) {
	tests := []struct {
		name string
		grid [][]int
		want int
	}{
		{"t1", [][]int{{3, 2, 1}, {1, 7, 6}, {2, 7, 7}}, 1},
		{"t2", [][]int{{3, 1, 2, 2}, {1, 4, 4, 5}, {2, 4, 2, 2}, {2, 4, 2, 2}}, 3},
		{"t3", [][]int{{13, 13}, {13, 13}}, 4},
		{"t4", [][]int{{11, 1}, {1, 11}}, 2},
		{"t5", [][]int{
			{3, 3, 3, 6, 18, 3, 3, 3, 3, 3},
			{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
			{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
			{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
			{1, 1, 1, 11, 19, 1, 1, 1, 1, 1},
			{3, 3, 3, 18, 19, 3, 3, 3, 3, 3},
			{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
			{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
			{3, 3, 3, 1, 6, 3, 3, 3, 3, 3},
			{3, 3, 3, 3, 1, 3, 3, 3, 3, 3},
		}, 48},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualPairsMethod01(tt.grid); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
