package _0240122

import (
	"testing"
)

var targetResult = make(map[int][][]int)

// https://leetcode.cn/problems/combination-sum/description/
// 无重复元素 的整数数组 candidates 和一个目标整数 target , 通过随意选取数组中数据求和得到tareget ,确保不同组合数少于 150 个
// 组合求和
// 思路：尝试构造递归实现，即将问题分解成子问题combinationSum(candidates, target-val)
func combinationSum(candidates []int, target int) [][]int {
	var rets [][]int
	if target == 0 {
		return nil
	}

	for _, num := range candidates {
		subTarget := target - num
		switch {
		case subTarget == 0: // 找了一个解
			rets = append(rets, []int{num})
		case subTarget > 0: // 还可以继续分
			// if vs, ok := targetResult[subTarget]; ok {
			// 	rets = append(rets, vs...)
			// } else {
			subRets := combinationSum(candidates, subTarget)
			for i, _ := range subRets {
				// check去重下
				ret := append(subRets[i], num)
				rets = append(rets, ret)
				// if !existRet(ret, subRets) {
				// }
			}
			// }
		case subTarget < 0: // 不合法
			continue
		}

	}

	// targetResult[target] = rets
	return rets
}

// 校对下是否存在
func existRet(input []int, rets [][]int) bool {
	toMap := func(arr []int) map[int]int {
		m := make(map[int]int)
		for _, v := range arr {
			m[v]++
		}
		return m
	}
	sameMap := func(m1, m2 map[int]int) bool {
		if len(m1) != len(m2) {
			return false
		}
		for k, _ := range m1 {
			if m1[k] != m2[k] {
				return false
			}
		}
		return true
	}
	m1 := toMap(input)
	for _, ret := range rets {
		if sameMap(m1, toMap(ret)) {
			return true
		}
	}
	return false
}

func TestCombinationSum(t *testing.T) {
	tests := []struct {
		name       string
		candidates []int
		target     int
		want       [][]int
	}{
		{"t1", []int{1}, 1, [][]int{{1}}},
		{"t2", []int{1}, 2, [][]int{{1, 1}}},
		{"t3", []int{1, 2}, 2, [][]int{{1, 1}, {2}}},
		{"t4", []int{2, 3, 5}, 5, [][]int{{2, 3}, {5}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := combinationSum(tt.candidates, tt.target)
			if len(got) != len(tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
