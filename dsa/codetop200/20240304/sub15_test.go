package _0240304

import (
	"testing"
)

func threeSum(nums []int) [][]int {
	var ret [][]int
	for i, num := range nums {
		idxs := twoSum(nums[i+1:], -num)
		if idxs != nil {
			ret = append(ret, []int{
				num,
				nums[i+idxs[0]],
				nums[i+idxs[1]],
			})
		}
	}
	return ret
}

func TestThreeSum(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want [][]int
	}{
		{"t1", []int{1, 1, -2, 3}, [][]int{{1, 1, -2}}},
		{"t2", []int{1, 1, -2, -2}, [][]int{{1, 1, -2}, {1, 1, -2}}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := threeSum(tt.nums)
			if len(got) != len(tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
