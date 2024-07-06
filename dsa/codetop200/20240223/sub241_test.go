package _0240223

import (
	"testing"
)

// 给定一个字符串，内容是数字字符，求该字符串表达的数值除以数值 n(n<10) 的结果，除不尽的保留 2 位小数（给出解题思路，并且代码实现）
// 例如: 给定字符串 "4562" 除以 5 的结果是 "912.4"
// 函数原型: char* DivString(char* dividend, int divisor);

// 思路:
//  1. 迭代字符串4562, ret=[9,1,2], concat = "4"
//  2. 45/5 = 9, 6/5=1.2,
func DivString(dividend string, divisor int32) string {

	// strconv.ParseInt()
	// var ret []string
	// var concat string
	// for _, c := range dividend {
	// 	// 4
	// 	num := c - '0'
	// 	if num < divisor { // 4<5
	// 		concat += fmt.Sprintf("%c", c)
	// 		continue
	// 	}
	// }
	//
	// k := struct {
	// 	a int
	// 	b int
	// }{100, 20}
	return ""
}

func TestLog(t *testing.T) {
	for _, c := range "4562" {
		t.Log(c - '0')
	}
}
