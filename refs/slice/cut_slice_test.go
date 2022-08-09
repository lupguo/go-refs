package slice

import (
	"math"
	"testing"
)

func runningSum(nums []int) []int {
	sumNums := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		sumNums[i] = sum(nums[0 : i+1])
	}

	return sumNums
}

// 计算nums的和
func sum(nums []int) int {
	var ret int
	for i := 0; i < len(nums); i++ {
		ret += nums[i]
	}
	return ret
}

func TestSliceCut(t *testing.T) {
	arr1 := []int{1, 2, 3, 4}
	for i := 0; i < len(arr1); i++ {
		t.Logf("arr1[0:%d]=%v", i, arr1[0:i+1])
	}
}

func TestCutSlice(t *testing.T) {
	var size = 25
	nums := make([]int, size)
	for i := 0; i < size; i++ {
		nums[i] = i
	}
	t.Logf("nums=>%+v", nums)

	perSize := 8
	count := int(math.Ceil(float64(len(nums)) / float64(perSize)))
	gnums := make([][]int, count)
	for i := 0; i < count; i++ {
		start := i * perSize
		end := (i + 1) * perSize
		if end > len(nums) {
			end = len(nums)
		}
		gnums[i] = nums[start:end]
		t.Logf("gnums[%d]=>%+v", i, gnums[i])
	}

}
