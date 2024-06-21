package leetcode

// 1. words[i] 和 words[j] 不包含相同字符
// 2. 长度乘积最大
func maxProduct(words []string) int {
	// 1. words 每个单词排序 + hash字符 [O(N)]
	// bword[i] = 0 1 0 1
	// bword[j] = 1 1 1 0
	bword := make([][]int, len(words))
	for i, word := range words {
		for j:=0; j<len(word); j++{
			bword[i][j] = word[j]
		}
	}


	// 2. 求不同 O(N^2), c[24]


	for i ,v := range c[24] {
		i := 0
		m := make(map[int]bool)
		for j, word := range words {
			if mark, ok := m[work[i]]; !ok {
				m[work[i]] = true
			}else if mark == true {

			}
			word[i] ==
		}
	}

	// 3. word[i] & word[j] == 0 => 不包含相同字符， 记录mul[i_j] = len(word[i])*len(word[j])

	// 非空 && 排序sort.Int(mul)，则排序取最大



}
