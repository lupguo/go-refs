package slice

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {
	tests := [][]string{
		{},
		{""},
	}
	for i, test := range tests {
		t.Logf("%v, %d=>%d", test, i, len(test))
	}
}

func TestCopySlice(t *testing.T) {
	r1 := []int{8, 9}
	r2 := &[]int{8, 9}
	r3 := &[]int{8, 9}
	op1 := func(k interface{}) {
		k = []int{1, 2, 3}
	}
	op2 := func(k interface{}) {
		k = &[]int{1, 2, 3}
	}
	op22 := func(k interface{}) {
		k = []int{1, 2, 3}
	}
	op1(r1)
	op2(r2)
	op22(r3)

	fmt.Printf("r1=%v, r2=%v, r3=%v", r1, r2, r3)
}
