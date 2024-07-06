package _0240304

import (
	"slices"
	"sort"
	"testing"
)

// 思路1: 将nums放到一个map中，m[val]=>i, 遍历m，检测是否存在m[target-val]在其中
// 这种无法应对存在重复元素到问题
func twoSum1(nums []int, target int) []int {
	m := make(map[int]int)
	for i, num := range nums {
		m[num] = i
	}

	for num, i := range m {
		if j, ok := m[target-num]; ok && i != j {
			return []int{i, j}
		}
	}
	return nil
}

type number struct {
	index int
	val   int
}

type numList []number

func (nums numList) Len() int {
	return len(nums)
}

func (nums numList) Less(i, j int) bool {
	return nums[i].val < nums[j].val
}

func (nums numList) Swap(i, j int) {
	nums[i], nums[j] = nums[j], nums[i]
}

// 思路2:
//  1. 先排序(构建排序+index值）
//  2. 利用left,right游标处理，在通过index获取元素在nums位置
func twoSum(nums []int, target int) []int {

	// 重组层numList，用于排序
	var list numList
	for i, num := range nums {
		list = append(list, number{
			index: i,
			val:   num,
		})
	}

	// 排序
	sort.Sort(list)

	// 左右游标靠拢
	left, right := 0, len(list)-1
	for left < right {
		numSum := list[left].val + list[right].val
		switch {
		case numSum > target: // 和太大，右边游标左移动
			right--
		case numSum < target: // 和太小，左边游标右移
			left++
		default:
			return []int{list[left].index, list[right].index}
		}
	}

	return nil
}

func TestTwoSum(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   []int
	}{
		{"t1", []int{1, 2, 3}, 3, []int{0, 1}},
		{"t2", []int{1, 2, 3}, 5, []int{1, 2}},
		{"t3", []int{1, 2, 3}, 6, nil},
		{"t4", []int{3, 3}, 6, []int{0, 1}},
		{"t5", []int{3, 2, 4}, 6, []int{1, 2}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := twoSum(tt.nums, tt.target)
			if !slices.Equal(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
