package leetcode

// https://leetcode.cn/leetbook/read/top-interview-questions-easy/x2y0c2/
// 给你两个整数数组 nums1 和 nums2 ，请你以数组形式返回两数组的交集。返回结果中每个元素出现的次数，应与元素在两个数组中都出现的次数一致（如果出现次数不一致，则考虑取较小值）。可以不考虑输出结果的顺序。
//

// 两个指针
func intersect(nums1 []int, nums2 []int) []int {
	m1 := make(map[int]int)
	for _, num := range nums1 {
		m1[num]++
	}

	//

	for _, num := range nums2 {
		if v, ok := m1[num]; ok {
			if v == 0 {

			}
		}
	}
}
