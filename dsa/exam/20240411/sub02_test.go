package _0240411

import (
	"math"
	"testing"

	"github.com/pkg/errors"
)

// 最大长度
const MaxStrLen = 11

// 字符串转整型
//
//	"4,578" -> 4578 = 4*10e3+5*10e2+7x10e1+8*10e0
//	"4578" ->
//	"-4" ->
//	思路: 依次迭代str, 转成数字字符，按位数进行处理
func strToInt(str string) (int, error) {
	// base case
	if str == "" {
		return 0, errors.New("err: empty string")
	}

	var chars []byte  // 数字切片 []rune{4,5,7,8}
	var negative bool // 符号位

	// 符号检测
	if str[0] == '-' {
		negative = true
		str = str[1:]
	} else if str[0] == '+' {
		str = str[1:]
	}

	// str[1:]检测
	if str == "" {
		return 0, errors.New("err: empty string")
	}

	// 数字切片
	for i := 0; i < len(str); i++ {
		c := str[i]
		switch {
		case c == ',': // 逗号跳过
			continue
		case c > '0' && c < '9': // 数字
			chars = append(chars, c)
		default:
			return 0, errors.Errorf("err: exception chars: %v", c) // 非,字符返回异常
		}
	}

	// 是否int越界
	n := len(chars) // []{4,5} = 4*10+5
	if n > MaxStrLen {
		return 0, errors.New("err: str number too long")
	}

	// 数字切片转整数
	var num int
	for i, c := range chars {
		num += int(c-'0') * int(math.Pow(10, float64(n-1-i)))
	}

	// 基于符号返回
	if negative {
		return -num, nil
	}

	return num, nil
}

func TestParse(t *testing.T) {
	
}

func TestStrToInt(t *testing.T) {
	tests := []struct {
		name string
		str  string
		want int
	}{
		{"t1", "4", 4},
		{"t2", "45", 45},
		{"t3", "4,578", 4578},
		{"t4", "4578", 4578},
		{"t5", "4578x", 0},
		{"t6", "-4578x", 0},
		{"t7", "-4", -4},
		{"t8", "04", 4},
		{"t9", "+4", 4},
		{"t10", "++4", 0},
		{"t11", ",4", 4},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, _ := strToInt(tt.str)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
