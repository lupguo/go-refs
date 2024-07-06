package buildinx

import (
	"reflect"
	"testing"
)

func TestCap(t *testing.T) {
	vals := map[string]any{
		"chan: no buff":              make(chan int),
		"chan: buff":                 make(chan int, 5),
		"slice: make([]int,0)":       make([]int, 0),
		"slice: make([]int,0,10)":    make([]int, 0, 10),
		"array: [100]int{} ":         [100]int{},
		"string: hello,中国":           "hello,中国",
		"map: make(map[int]int, 20)": make(map[int]int, 20),
	}

	for k, v := range vals {
		refVal := reflect.ValueOf(v)
		t.Logf("%v len(v)=> %v, cap(v) => %v ", k, refVal.Len(), refVal.Cap())
	}
}

func TestCapChan(t *testing.T) {
	ch := make(chan int, 3)
	ch <- 1
	t.Logf("等待接收:%v", len(ch)) // 输出：1
	ch <- 2
	t.Logf("等待接收:%v", len(ch)) // 输出：2
	t.Logf("容量:%v", cap(ch))   // 输出：3
}
