package bitmap

import (
	"fmt"
	"testing"
)

func TestNewBitmap(t *testing.T) {
	bm1 := NewBitmap(1)
	bm1.Set(3)
	t.Log(bm1.Get(3))
	t.Log(bm1.Get(2))
}

func Test_intToBinary(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"t1", args{1}, fmt.Sprintf("%032b", 1)},
		{"t2", args{3}, fmt.Sprintf("%032b", 3)},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := intToBinary(tt.args.num)
			t.Log(got)
			if got != tt.want {
				t.Errorf("intToBinary() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_hammingWeight(t *testing.T) {
	type args struct {
		num uint32
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{num: 3}, 2},
		{"t2", args{num: 4}, 1},
		{"t2", args{num: 8}, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingWeight(tt.args.num); got != tt.want {
				t.Errorf("hammingWeight() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_countOnes(t *testing.T) {
	type args struct {
		num int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{"t1", args{num: 3}, 2},
		{"t1", args{num: 4}, 1},
		{"t1", args{num: 5}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countOnes(tt.args.num); got != tt.want {
				t.Errorf("countOnes() = %v, want %v", got, tt.want)
			}
		})
	}
}
