package slice

import (
	"math"
	"testing"
)

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
