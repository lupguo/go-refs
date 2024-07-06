package week12

import (
	"testing"
)

func Test_halveArray(t *testing.T) {

	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"t1", []int{1}, 1},
		{"t2", []int{1, 2}, 2},
		{"t3", []int{1, 2, 3}, 2},
		{"t4", []int{5, 19, 8, 1}, 3},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := halveArray(tt.nums); got != tt.want {
				t.Errorf("halveArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
