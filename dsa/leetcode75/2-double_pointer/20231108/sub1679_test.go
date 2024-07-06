package _0231108

import (
	"slices"
	"testing"
)

// 给你一个整数数组 nums 和一个整数 k 。
//
// 每一步操作中，你需要从数组中选出和为 k 的两个整数，并将它们移出数组。
//
// 返回你可以对数组执行的最大操作数。
//
// 输入：nums = [1,2,3,4], k = 5
// 输出：2
// 解释：开始时 nums = [1,2,3,4]：
// - 移出 1 和 4 ，之后 nums = [2,3]
// - 移出 2 和 3 ，之后 nums = []
// 不再有和为 5 的数对，因此最多执行 2 次操作。
//
// 示例 2：
//
// 输入：nums = [3,1,3,4,3], k = 6
// 输出：1
// 解释：开始时 nums = [3,1,3,4,3]：
// - 移出前两个 3 ，之后nums = [1,4,3]
// 不再有和为 6 的数对，因此最多执行 1 次操作。
func maxOperations(nums []int, k int) int {

	return 0
}

// 方法4：map的优化版本，初始map同时就开始检测是否存在map[k-num]存在，如果存在则找到目标值
func Method4MaxOperations(nums []int, k int) int {
	numCntMap := make(map[int]int)

	// 初始化构造一个map，将Num作为map的key，同时针对key计数
	var found int
	for _, num := range nums {
		numCntMap[num]++
		target := k - num
		if numCntMap[target] > 0 { // 找到了一个，不过应该注意num
			// 如果num和target相等，但不满足numCntMap[target] >=2 的要求，则不符合
			if target == num && numCntMap[target] < 2 {
				continue
			}

			numCntMap[num]--
			numCntMap[target]--
			found++
		}
	}

	return found
}

// 方法3: 尝试使用hash map，遍历初始化数组到map, 检测m[k-nums[i]]是否存在
// 不过因为在nums中可能存在两个值一样，可以考虑将map设置键位数字，值为可用次数，如果为0则表示不可用了
// 时间复杂度O(N)，空间复杂度O(N)存储map
func Method3MaxOperations(nums []int, k int) int {
	numCntMap := make(map[int]int)

	// 初始化构造一个map，将Num作为map的key，同时针对key计数
	for _, num := range nums {
		numCntMap[num]++
	}

	// 尝试迭代nums，从map中找k-num的，如果找到则将对应的计数扣除
	// p.s注意即便num没有找到，也不用重新加回来，因为这个值如果在map中如果找不到目标值，该值扣减到0也无影响
	var found int
	for _, num := range nums {
		// 如果num已经没有可用的计数了，直接跳过num（可能在前面匹配时候已经把numCntMap[num]占用扣减掉了）
		if numCntMap[num] == 0 {
			continue
		}

		// num还有可用的计数
		numCntMap[num]--
		target := k - num
		if numCntMap[target] > 0 {
			found++
			numCntMap[target]--
		}
	}

	return found
}

// 方法2：先排序，在双指针i,j 两边夹逼处理 i=0, j=len-1
// 1. 先排序，有序后找值
// 2. 思考: 如果小数num[i]+大数num[j] > k，此时只有将j左移(j--)，才能找到可能符合=k的解
// 3. 同理，如果小数num[i]+大数num[j] < k, 此时只有将i右移(i++)，才能找到可能符合=k的解
// 如果num[i]+num[j] = k, 则i++, j--，并统计found++
func Method2MaxOperations(nums []int, k int) int {
	// 排序 O(NLogN)
	slices.Sort(nums)

	// 双边夹逼 O(N)
	var found int
	n := len(nums)
	i, j := 0, n-1
	for i < j {
		sum := nums[i] + nums[j]
		switch {
		case sum > k: // 两者和过大，将大数位置j左移
			j--
		case sum < k: // 两者和过小，将小数位置j右移
			i++
		default: // 找到目标值，同时移动i,j
			i++
			j--
			found++
		}
	}

	return found
}

// 方法1： 暴力算法
// 1. i,j 两个指针，i+j=k，配置一个map用于记录i,j是否被使用
// 2. 尝试剪枝: nums[i] >k or nums[j] > k，直接标记失效
// 算法复杂度: O(N^2)
func Method1MaxOperations(nums []int, k int) int {

	var found int
	used := make(map[int]bool)
	over := make(map[int]bool)
	for i := 0; i < len(nums); i++ {
		if nums[i] > k {
			over[i] = true
		}

		// 如果i用过了 或者过大，跳过
		if used[i] || over[i] {
			continue
		}

		for j := i + 1; j < len(nums); j++ {
			if nums[j] > k {
				over[j] = true
			}

			// 如果j用过了，或者过大，跳过
			if used[j] || over[j] {
				continue
			}

			// 找到一个目标值, 并将它们移出数组，结束本轮循环
			if nums[i]+nums[j] == k {
				used[i] = true
				used[j] = true
				found++

				// 这里要break出来，num[i]已被用了
				break
			}
		}

	}

	return found
}

func TestMethod1MaxOperations(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"t1", []int{1}, 1, 0},
		{"t2", []int{1, 2, 3, 4}, 5, 2},
		{"t3", []int{3, 1, 3, 4, 3}, 6, 1},
		{"t4", []int{4, 4, 1, 3, 1, 3, 2, 2, 5, 5, 1, 5, 2, 1, 2, 3, 5, 4}, 2, 2},
		{"t5", []int{2, 5, 4, 4, 1, 3, 4, 4, 1, 4, 4, 1, 2, 1, 2, 2, 3, 2, 4, 2}, 3, 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Method4MaxOperations(tt.nums, tt.k); got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
