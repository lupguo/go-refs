package _0240306

// 有序数组，两数之和
//
//	https://leetcode.cn/problems/kLl5u1/description/
//
// 思路:
//  1. 有序数组
//  2. i,j双边求和，边界收缩求解
//  3. 只存在一对
func twoSum(numbers []int, target int) []int {
	i, j := 0, len(numbers)-1
	for i < j {
		sumTwo := numbers[i] + numbers[j]
		if sumTwo == target {
			return []int{i, j}
		} else if sumTwo > target { // 双数和太大，有序递增，右边界左移
			j--
		} else { // 双数和太小，左边界右移
			i++
		}
	}

	return nil
}
