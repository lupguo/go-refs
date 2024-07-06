package _0240311

// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。
// 时间复杂度为 O(n) 的算法
// 思路1:
//  1. sort(nums) -> nums[k] ONLogN
//  2. 堆排 -> 构建K个元素小顶堆，依次将nums元素num > 堆顶，则加入堆中做堆化处理，依次小顶堆包含了nums的k个最大元素，且堆顶为这个k个最大元素最小值，即要得到到解
//  2. 快排分区法
//  1. 求k大元素，即排序后位置为k'=n-k索引的元素
//  2. 尝试通过分区函数p = partition(nums, low, high)
//     若p < k', 则 k'在分区点右侧，则low=p+1
//     若p > k', 则 k'在分区点左侧，则high=p-1
//     若p == k'，则找到第k大元素，返回值
//     2.3 返回pivot值即第k大元素
func findKthLargest(nums []int, k int) int {
	// 基础条件
	n := len(nums)
	if k > n || nums == nil {
		return 0
	}

	// 不断在low, high之前寻求p，看p是否是要找到第k大元素位置
	wantP := n - k
	low, high := 0, n-1
	for {
		p := partition(nums, low, high) // 分区[
		switch {
		case p < wantP: // 分区点位置靠左
			low = p + 1
		case p > wantP: // 分区点位置靠右，继续在左侧找
			high = p - 1
		default:
			return nums[p]
		}
	}
}

// 选取最后元素为pivot分区点进行分区，得到首个pivot元素位置
func partition(nums []int, low int, high int) int {
	if len(nums) < 2 {
		return 0
	}

	// 设定最高点为分区点
	pivot := nums[high]

	// i,j 从左到右，nums[0:i] < pivot < nums[j:]
	i, j := low, low
	for j < high {

		// 探测j找到一个比pivot小的，需要和nums[i]交换，
		// 4,   2,  1, [3]
		// i,j
		// i,   j (swap)
		// 2,   4,  1, [3]
		//      i,  j(swap)
		// 2,   1,  4,  [3]
		//          i,  j(break for)
		//          i<->high(swap)
		// 2,   1,  [3],  4
		//          done
		if nums[j] < pivot {
			nums[i], nums[j] = nums[j], nums[i]
			i++
		}
		j++
	}

	// 交换pivot和nums[i]
	nums[i], nums[high] = nums[high], nums[i]

	return i
}
