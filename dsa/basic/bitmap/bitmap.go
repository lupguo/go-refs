package bitmap

import (
	"fmt"
	"strings"
)

type Bitmap struct {
	bits []int
}

func NewBitmap(n int) *Bitmap {
	return &Bitmap{
		bits: make([]int, (n+31)/32),
	}
}

func (b *Bitmap) Set(k uint) {
	idx := k / 32
	pos := k % 32
	fmt.Printf("idx=%v, pos=%v, 1<<%[2]v = %08b\n", idx, pos, 1<<pos)
	fmt.Printf("b.bits[idx] =%08b | 1<<pos = %08b\n", b.bits[idx], 1<<pos)
	b.bits[idx] |= 1 << pos
}

func (b *Bitmap) Get(k uint) int {
	idx := k / 32
	pos := k % 32
	return (b.bits[idx] >> pos) & 1
}

// 要求给定一个整数，输出在其二进制表示中1的个数。
//
// 示例：
// 输入: 00000000000000000000000000001011
// 输出: 3
// 解释: 输入的二进制表示为 00000000000000000000000000001011 ，其中有三位为 '1'。
// 解题思路：可以不断将该整数的最后一位1去除，并累加计数器
func hammingWeight(num uint32) int {
	cnt := 0
	for num != 0 {
		cnt++
		num &= num - 1
	}
	return cnt
}

// 不断将整数除以2来模拟去掉最后一位1的操作，累加计数器，直到整数变为0为止，这样得到的计数器值就是原整数的二进制表示中1的个数。
func countOnes(num int) int {
	count := 0
	for num != 0 {
		if num&1 == 1 { // 最后一位是否为1，若为1则计数器加一
			count++
		}
		num >>= 1 // 右移迭代
	}
	return count
}

// 将整数不断右移，并判断其二进制最后一位是否为1
func intToBinary(num int) string {
	var ret []string
	for i := 31; i >= 0; i-- {
		if (num & (1 << i)) != 0 {
			ret = append(ret, "1")
		} else {
			ret = append(ret, "0")
		}
	}

	return strings.Join(ret, "")
}
