package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func addBinary2(a string, b string) string {
	parseIntA, err := strconv.ParseInt(a, 2, 0)
	if err != nil {
		return err.Error()
	}
	parseIntB, err := strconv.ParseInt(b, 2, 0)
	if err != nil {
		return err.Error()
	}
	return fmt.Sprintf("%b", parseIntA+parseIntB)
}

func addBinary(a string, b string) string {
	// 字符串转int切片
	sa := toReverseInts(a)
	sb := toReverseInts(b)

	la, lb := len(sa), len(sb)
	lc := max(la, lb) + 1
	sc := make([]int, lc)

	addBase := 0 // 进制记录
	for i := 0; i < lc; i++ {
		if i < la && i < lb {
			sc[i] = sa[i] + sb[i] + addBase
		} else if i >= la {
			if i < lb {
				sc[i] = sb[i] + addBase
			} else {
				sc[i] = addBase
			}
		} else if i >= lb {
			if i < la {
				sc[i] = sa[i] + addBase
			}
		}

		sc[i] = sc[i] % 2
		addBase = sc[i] / 2
	}

	// int切片转字符串
	return toStr(sc)
}

// 较大值
func max(i, j int) int {
	if i >= j {
		return i
	}
	return j
}

// 转成整型切片
func toReverseInts(ss string) []int {
	strings.Split(ss, "")
	var ret []int
	for i := len(ss) - 1; i >= 0; i-- {
		if ss[i] == '1' {
			ret = append(ret, 1)
		} else {
			ret = append(ret, 0)
		}
	}
	return ret
}

func toStr(vv []int) string {
	var s string
	for i := len(vv) - 1; i >= 0; i-- {
		if s == "" && vv[i] == 0 {
			continue
		}
		s = fmt.Sprintf("%s%d", s, vv[i])
	}
	if s == "" {
		return "0"
	}
	return s
}
