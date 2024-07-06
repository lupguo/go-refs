package _0240125

import (
	"fmt"
	"testing"
)

func TestSliceAppend(t *testing.T) {
	que := []int{1, 2, 3}
	i := 100
	lastCap := cap(que)
	for len(que) > 0 {
		// 出队
		// val := que[0]
		// que = que[1:]

		// 模拟队列扩展
		que = append(que, i)
		i++
		if lastCap != cap(que) {
			fmt.Printf("que cap=[%v], que len=[%v], node val=>%v\n", cap(que), len(que), 'x')
			lastCap = cap(que)
		}

		if i > 1e4 {
			break
		}
	}

	// t.Logf("out que %v", que)
}
