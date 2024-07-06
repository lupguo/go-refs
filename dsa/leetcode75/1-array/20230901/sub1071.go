package _0230901

// 解题思路
// 1. 寻找单个字符串的可能存在的公因字符串，以切片返回
// 2. 寻找两个GCD切片中，最长的GCD

// https://leetcode.cn/problems/greatest-common-divisor-of-strings/?envType=study-plan-v2&envId=leetcode-75
//
// 对于字符串 s 和 t，只有在 s = t + ... + t（t 自身连接 1 次或多次）时，我们才认定 “t 能除尽 s”。
// 给定两个字符串 str1 和 str2 。返回 最长字符串 x，要求满足 x 能除尽 str1 且 x 能除尽 str2 。
// 输入：str1 = "ABCABC", str2 = "ABC"
// 输出："ABC"
func gcdOfStrings(str1 string, str2 string) string {
	str1Gcds, str2Gcds := gcdListOfString(str1), gcdListOfString(str2)

	return maxLenGcd(str1Gcds, str2Gcds)
}

func maxLenGcd(gcds1 []string, gcds2 []string) string {
	len1, len2 := len(gcds1), len(gcds2)
	var maxLenGcd string
	for i, j := 0, 0; i < len1 && j < len2; {
		li, lj := len(gcds1[i]), len(gcds2[j])
		if li > lj {
			j++
		} else if li < lj {
			i++
		} else { // 长度相等，检测是否gcd一致
			if gcds1[i] == gcds2[j] {
				maxLenGcd = gcds1[i]
			}
			i++
			j++
		}
	}

	return maxLenGcd
}

// 公共列表
func gcdListOfString(str string) []string {
	if str == "" {
		return []string{}
	}
	var gcds []string
	strLen := len(str)
	for i := 1; i < strLen/2+1; i++ {
		gcd := str[:i]
		gcdLen := len(gcd)

		// 是否gcd长度能被除尽，如果不能，则该gcd肯定不正确，换下个
		if strLen%gcdLen != 0 {
			continue
		}

		// 检测gcd是否正确，从起始位依次开始校验，检测每个位置是否和原位置字符串子串一致
		found := true
		for k := 0; k < strLen; k = k + gcdLen {
			if str[k:k+gcdLen] != gcd {
				found = false
				break
			}
		}
		if found {
			gcds = append(gcds, gcd)
		}
	}

	// 自己字符串也需要算上
	gcds = append(gcds, str)

	return gcds
}
