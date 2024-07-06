package _0240312

// 给定一个未排序的整数数组 nums ，找出数字连续的最长序列（不要求序列元素在原数组中连续）的长度。
//
// 请你设计并实现时间复杂度为 O(n) 的算法解决此问题。
// 输入：nums = [100,4,200,1,3,2]
// 输出：4
// 解释：最长数字连续序列是 [1, 2, 3, 4]。它的长度为 4。

// 思路:
// 1. 桶排序，找到数组最大值，申请对应空间，放到对应位置空间(O(max(nums))、时间O(N) -> 解有问题（当为负数时候无法使用下标、如何区分0）
// 1. 改成map，将数组值存到map中，再寻找map的首个最小值，再依次迭代map找到最大值？
// 2. i迭代获取最大位置，更替maxLen、O(N)
func longestConsecutive(nums []int) int {
	if nums == nil {
		return 0
	}

	// 转存到map中
	numMap := make(map[int]bool)
	for _, num := range nums {
		numMap[num] = true
	}

	// 获取map的首个元素位置
	var maxLen int
	for num, _ := range numMap {
		// 查找到当前num元素所在序列的首个位置
		if numMap[num-1] {
			continue
		}

		// 找到最小元素
		minNum := num

		// 从minNum元素的开始位置迭代，依次迭代更替长度
		var cnt int
		for numMap[minNum] {
			cnt++
			minNum++
		}
		maxLen = max(maxLen, cnt)

	}

	return maxLen
}
