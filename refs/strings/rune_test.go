package strings

import (
	"testing"
)

func TestRune(t *testing.T) {
	// var (
	// 	a1 = 1      // 10进制int整型
	// 	a2 = 01     // 8进制数字1
	// 	a3 = 0x1    // 16进制数字1
	// 	a4 = 0b0001 // 2进制数字1
	//
	// 	b1 = '1' // 字符
	// 	b2 = "1" // 字符串1
	// 	b3 = `1` // 字面量1
	// )

	numbers := []interface{}{
		1, 01, 0x1, 0b0001,
	}

	strs := []interface{}{
		'1', "1", `1`,
		'中', "中国", `中国`,
	}

	strs2 := "Go语言编程言简意赅"

	for i, number := range numbers {
		t.Logf("i[%v]: %v, %#v, %+v", i, number, number, number)
	}
	for i, str := range strs {
		t.Logf("i[%v]: %v, %#v, %+v", i, str, str, str)
	}
	for i, r := range strs2 {
		t.Logf("i[%d]: %c, 2进制[%b], 16进制[0x%0x], 类型T[%T], %v, %#v, %+v", i, r, r, r, r, r, r, r)
	}
	for i, r := range []byte(strs2) {
		t.Logf("i[%d]: %c, 2进制[%b], 16进制[0x%0x], 类型T[%T],%v, %#v, %+v", i, r, r, r, r, r, r, r)
	}
	for i, r := range []rune(strs2) { // 实际和range 字符串处理结果一致
		t.Logf("i[%d]: %c, 2进制[%b], 16进制[0x%0x], 类型T[%T],%v, %#v, %+v", i, r, r, r, r, r, r, r)
	}
}

func TestStringsLength(t *testing.T) {
	str1 := "言简意赅"
	t.Logf("len(str),返回字节数=%d", len(str1))
	t.Logf("len([]byte(str)),返回字节数(utf8编码)，和len(str)是一致=%d", len([]byte(str1)))
	t.Logf("len([]rune(str)), 返回字面量数=%d", len([]rune(str1)))
}
