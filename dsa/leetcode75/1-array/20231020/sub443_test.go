package _0231020

import (
	"reflect"
	"strconv"
	"testing"
)

// 给你一个字符数组 chars ，请使用下述算法压缩：
//
// 从一个空字符串 s 开始。对于 chars 中的每组 连续重复字符 ：
//   - 如果这一组长度为 1 ，则将字符追加到 s 中。 - 用例1
//   - 否则，需要向 s 追加字符，后跟这一组的长度(注意: 如果组长度为 10 或 10 以上，则在 chars 数组中会被拆分为多个字符) - 用例2
//
// 示例 1：
//
// 输入：chars = ["a","a","b","b","c","c","c"]
// 输出：返回 6 ，输入数组的前 6 个字符应该是：["a","2","b","2","c","3"]
// 解释："aa" 被 "a2" 替代。"bb" 被 "b2" 替代。"ccc" 被 "c3" 替代。
// 示例 2：
//
// 输入：chars = ["a"]
// 输出：返回 1 ，输入数组的前 1 个字符应该是：["a"]
// 解释：唯一的组是“a”，它保持未压缩，因为它是一个字符。
// 示例 3：
//
// 输入：chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
// 输出：返回 4 ，输入数组的前 4 个字符应该是：["a","b","1","2"]。
// 解释：由于字符 "a" 不重复，所以不会被压缩。"bbbbbbbbbbbb" 被 “b12” 替代。
//
// 压缩后得到的字符串 s 不应该直接返回 ，需要转储到字符数组 chars 中; '
//
// 请在 修改完输入数组后 ，返回该数组的新长度。
func compress(chars []byte) int {
	chars, n := compressChars(chars)
	return n
}

func compressChars(chars []byte) ([]byte, int) {
	if len(chars) == 0 {
		return nil, 0
	}

	var wantChars []byte
	curChar := chars[0]
	curCnt := 0
	length := len(chars)
	for i, c := range chars {
		// 当前迭代的c和前一个没有变化，则计数加1
		if curChar == c {
			curCnt++

			// 如果迭代到最后一个元素，则直接将当前字符计数加入到结果切片中
			if i == length-1 {
				wantChars = append(wantChars, toChars(curChar, curCnt)...)
			}

			continue
		}

		// c有变化了，计数完成，记录curChar的计数，并更新curChar,curCnt
		wantChars = append(wantChars, toChars(curChar, curCnt)...)
		curChar, curCnt = c, 1 // 更新curChar和curCnt
	}

	return wantChars, len(wantChars)
}

// 组合char和char的数量成字符切片
func toChars(char byte, cnt int) []byte {
	var charCounts []byte
	charCounts = append(charCounts, char)
	if cnt < 2 {
		return charCounts
	}

	charCounts = append(charCounts, strconv.Itoa(cnt)...)

	return charCounts
}

func TestToChars(t *testing.T) {
	tests := []struct {
		name string
		char byte
		cnt  int
		want []byte
	}{
		{"t1", 'a', 10, []byte{'a', '1', '0'}},
		{"t2", 'a', 1, []byte{'a'}},
		{"t3", 'a', 5, []byte{'a', '5'}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := toChars(tt.char, tt.cnt); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}

}

func TestCompress(t *testing.T) {
	tests := []struct {
		name      string
		chars     []byte
		wantChars []byte
		want      int
	}{
		{"t1", []byte{'a', 'a', 'b', 'b', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '2'}, 6},
		{"t2", []byte{'a', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b', 'b'}, []byte{'a', 'b', '1', '2'}, 4},
		{"t3", []byte{'a'}, []byte{'a'}, 1},
		{"t4", []byte{'a', 'a', 'b', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', '2', 'c', '3'}, 6},
		{"t5", []byte{'a', 'a', 'b', 'c', 'c', 'c'}, []byte{'a', '2', 'b', 'c', '3'}, 5},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotChars, n := compressChars(tt.chars)
			if !reflect.DeepEqual(gotChars, tt.wantChars) {
				t.Errorf("input:%[1]s, gotChars: %v, %[2]s, but want: %v, %[3]s", tt.chars, gotChars, tt.wantChars)
			}
			if n != tt.want {
				t.Errorf("gotN: %v, but want: %v", n, tt.want)
			}
		})
	}
}
