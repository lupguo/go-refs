package _0240124

import (
	"fmt"
	"reflect"
	"testing"
)

// 快排
func sortArray(nums []int) []int {
	return quickSort(nums, 0, len(nums)-1)
}

// 快排
//
//	思路: 分区 partition、两边交换
func quickSort(nums []int, low, high int) []int {
	if len(nums) <= 1 || low < 0 || high < 0 {
		return nums
	}
	p := partition(nums, low, high)

	quickSort(nums, low, p-1)
	quickSort(nums, p+1, high)

	return nums
}

func TestQuickSort(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		output []int
	}{
		{"t1", []int{3, 2, 1, 5, 6, 4}, []int{1, 2, 3, 4, 5, 6}},
		{"t2", []int{1, 2, 3, 4, 5, 6}, []int{1, 2, 3, 4, 5, 6}},
		{"t3", []int{6, 5, 4, 3, 2, 1}, []int{1, 2, 3, 4, 5, 6}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := quickSort(tt.nums, 0, len(tt.nums)-1)

			if !reflect.DeepEqual(got, tt.output) {
				t.Errorf("got %v, but want %v", got, tt.output)
			}
		})
	}
}

// 找到分区点，并将排好pivot位置的值 -  Lomuto分区方案
func partitionLomuto(nums []int, lo, hi int) int {
	// https://en.wikipedia.org/wiki/Quicksort
	pivot := nums[hi]          // 用最尾数做pivot，选择最后一个位置作为pivot临时参考值
	pivotIndex := lo           // pivotIndex 存储着大于pivot值的index（和nums[i] <= pivot相关，如果是nums[i] < pivot，则pivot存储是不小于pivot的值节点）
	for i := lo; i < hi; i++ { // i 是快指针,做节点探测, 如果找到 >= piv
		if nums[i] <= pivot { //
			nums[pivotIndex], nums[i] = nums[i], nums[pivotIndex] // 因为pivotIndex默认是存在着 “大于”pivot的值，这里先交换
			pivotIndex++                                          //
		}
	}
	// 交互，将pivot和pivotIndex匹配到正确的位置
	nums[pivotIndex], nums[hi] = nums[hi], nums[pivotIndex]

	return pivotIndex
}

// Hoare 排序方案 (左、右指针，向中间靠拢） - 减少了交换次数
func partition(nums []int, lo, hi int) int {
	pivot := nums[hi]
	left, right := lo, hi
	var round int
	for left < right {
		round++
		fmt.Printf("round(%v), nums=%v\n", round, nums)
		fmt.Printf("round(%v), start: pivot=%v, l=%v(val=%v), r=%v(val=%v)\n", round, pivot, left, nums[left], right, nums[right])
		// 从左向右选择找大于pivot的节点，找到退出循环
		for left < right && nums[left] < pivot {
			fmt.Printf("round(%v), l++\n", round)
			left++
		}
		fmt.Printf("round(%v), after left move ->: pivot=%v, l=%v(val=%v), r=%v(val=%v)\n", round, pivot, left, nums[left], right, nums[right])
		// 从右向向左寻找小于pivot节点，找到退出循环
		for left < right && nums[right] > pivot {
			fmt.Printf("round(%v), r--\n", round)
			right--
		}
		fmt.Printf("round(%v), after right move <-: pivot=%v, l=%v(val=%v), r=%v(val=%v)\n", round, pivot, left, nums[left], right, nums[right])

		// swap or break
		if nums[left] == nums[right] {
			break
		}
		nums[left], nums[right] = nums[right], nums[left]
		fmt.Printf("round(%v), after swap, nums=%v\n\n", round, nums)
	}

	fmt.Printf("fin, partition=%v(%v)\n\n", left, nums[left])
	return left
}

func TestPartition(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{3, 2, 1, 5, 6, 4}, 3},
		{"t2", []int{1, 2, 3, 4, 5, 6}, 5},
		{"t3", []int{6, 5, 4, 3, 2, 1}, 0},
		{"t4", []int{1, 2, 3}, 2},
		{"t5", []int{1, 5, 3, 2, 4}, 3},
		{"t5-重复元素", []int{1, 2, 3, 2, 2}, 1},
		{"t6-重复元素", []int{3, 2, 2}, 0},
		{"t6", []int{4, 1, 2, 3}, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := partition(tt.nums, 0, len(tt.nums)-1)
			t.Logf("tt.nums=%v", tt.nums)

			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}

func TestSwapSlice(t *testing.T) {
	nums := []int{1, 2, 3}
	swapNums(nums)
	t.Log(nums)
	nums2 := appendNums(&nums, []int{100, 200, 300})
	t.Log(nums, nums2)
}

func appendNums(nums *[]int, ints []int) []int {
	*nums = append(*nums, ints...) // nums进行了扩容，如果想在函数外同步，需要通过指针应用方式修改
	return *nums
}

// 因为slice是copy了引用，在函数内swap会反馈到函数外
// 但如果append进行了扩容，情况又不一样了
func swapNums(nums []int) {
	n := len(nums)
	nums[0], nums[n-1] = nums[n-1], nums[0]
}
