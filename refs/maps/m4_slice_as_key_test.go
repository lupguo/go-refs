package maps

import (
	"testing"
)

func TestSliceAsKey(t *testing.T) {
	s1 := [][]int{{1, 2}, {3}, {4, 5, 6}}
	m1 := make(map[interface{}]int)
	for i, v := range s1 {
		t.Logf("#%d=>%v, ", i, v)
		m1[v] = len(v)
	}

	for k, v := range m1 {
		t.Logf("key:%v(%T),val:%v", k, k, v)
	}
}

func TestStructAsKey(t *testing.T) {
	type user struct {
		ID   int
		Name string
	}
	// 结构体作为key
	m1 := make(map[user]int)
	m1[user{1, "terry"}] = 100
	m1[user{2, "clark"}] = 200

	for k, v := range m1 {
		t.Logf("key:%v(%T),val:%v", k, k, v)
	}
}

func TestArrayAsKey(t *testing.T) {
	a1 := [3][3]int{{1, 2}, {3}, {4, 5, 6}}

	// 用interface{}作为key
	m1 := make(map[interface{}]int)

	// 数组求和
	sum := func(arr [3]int) int {
		s := 0
		for _, v := range arr {
			s += v
		}
		return s
	}

	// 赋值map
	for i, v := range a1 {
		t.Logf("#%d=>%v, ", i, v)
		m1[v] = sum(v)
	}

	// 输出结果
	for k, v := range m1 {
		t.Logf("key:%v(%T),val:%v", k, k, v)
	}
}
