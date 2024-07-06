package _0230901

import (
	"reflect"
	"testing"
)

// https://leetcode.cn/problems/kids-with-the-greatest-number-of-candies/description/?envType=study-plan-v2&envId=leetcode-75
//
// 给你一个数组 candies 和一个整数 extraCandies ，其中 candies[i] 代表第 i 个孩子拥有的糖果数目。
//
// 对每一个孩子，检查是否存在一种方案，将额外的 extraCandies 个糖果分配给孩子们之后，此孩子有 最多 的糖果。注意，允许有多个孩子同时拥有 最多 的糖果数目。
//
// 示例 1：
//
// 输入：candies = [2,3,5,1,3], extraCandies = 3
// 输出：[true,true,true,false,true]
// 解释：
// 孩子 1 有 2 个糖果，如果他得到所有额外的糖果（3个），那么他总共有 5 个糖果，他将成为拥有最多糖果的孩子。
// 孩子 2 有 3 个糖果，如果他得到至少 2 个额外糖果，那么他将成为拥有最多糖果的孩子。
// 孩子 3 有 5 个糖果，他已经是拥有最多糖果的孩子。
// 孩子 4 有 1 个糖果，即使他得到所有额外的糖果，他也只有 4 个糖果，无法成为拥有糖果最多的孩子。
// 孩子 5 有 3 个糖果，如果他得到至少 2 个额外糖果，那么他将成为拥有最多糖果的孩子。
func kidsWithCandies(candies []int, extraCandies int) []bool {
	var ret []bool
	if len(candies) == 0 {
		return ret
	}

	// 找到拥有糖果最多的孩子
	var maxCnt int
	for _, cnt := range candies {
		if cnt > maxCnt {
			maxCnt = cnt
		}
	}

	// 假定重新让每个孩子分配额外糖果后，看是否达到拥有最多糖果孩子要求
	for _, cnt := range candies {
		if cnt+extraCandies >= maxCnt {
			ret = append(ret, true)
		} else {
			ret = append(ret, false)
		}
	}

	return ret
}

func TestKidsWithCandies(t *testing.T) {
	candies := []int{2, 3, 5, 1, 3}
	extraCandies := 3
	want := []bool{true, true, true, false, true}
	got := kidsWithCandies(candies, extraCandies)

	if reflect.DeepEqual(want, got) == false {
		t.Errorf("want=%v, but got %v", want, got)
	}
}
