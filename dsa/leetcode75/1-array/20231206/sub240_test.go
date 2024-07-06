package _0231206

import (
	"testing"
)

// 编写一个高效的算法来搜索 m x n 矩阵 matrix 中的一个目标值 target 。
//
// 每行的元素从左到右升序排列。
// 每列的元素从上到下升序排列。
//
// 思路:  基于二维数组存在行列的递增特性，考虑从右上角->左下角递进寻找目标值
//  1. 异常判断, matrix == nil
//  2. 得到行列值: row=len(matrix), col=len(matrix[0])
//  3. 循环列表: for j:=col-1; j >=0; j-- {}
//     循环行: for i:=0; i<row; i++ {
//     if target == matrix[i][j] return
//     }
//     时间复杂度O(m+n) 空间复杂度O(1)
func searchMatrix(matrix [][]int, target int) bool {
	if matrix == nil {
		return false
	}

	// 获取行列值
	row, col := len(matrix), len(matrix[0])
	for j := col - 1; j >= 0; j-- {
		for i := 0; i < row; i++ {
			if target == matrix[i][j] {
				return true
			}
		}
	}

	return false
}

func TestSearchMatrix(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		target int
		want   bool
	}{
		{"t1", [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 5, true},
		{"t2", [][]int{
			{1, 4, 7, 11, 15},
			{2, 5, 8, 12, 19},
			{3, 6, 9, 16, 22},
			{10, 13, 14, 17, 24},
			{18, 21, 23, 26, 30},
		}, 20, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := searchMatrix(tt.matrix, tt.target); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
