package _0240314

import (
	"testing"
)

func Test_trap(t *testing.T) {
	tests := []struct {
		name   string
		height []int
		want   int
	}{
		{"t1", []int{0, 1}, 0},
		{"t2", []int{0, 1, 0, 1}, 1},
		{"t3", []int{1, 1, 0, 1}, 1},
		{"t4", []int{1, 0, 0, 1}, 2},
		{"t5", []int{1, 0, 0, 2}, 2},
		{"t6", []int{1, 0, 1, 2}, 1},
		{"t7", []int{1, 0, 1, 2, 3}, 1},
		{"t8", []int{1, 0, 0, 2, 3}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := trap(tt.height)
			if got != tt.want {
				t.Errorf("got %v, but want %v", got, tt.want)
			}
		})
	}
}
