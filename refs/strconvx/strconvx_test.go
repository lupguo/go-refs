package strconvx

import (
	"strconv"
	"testing"
)

func TestStrConvInt(t *testing.T) {
	groupID := "485635"

	type args struct {
		base    int
		bizSize int
	}

	tests := []struct {
		name string
		args args
	}{
		{"10_0", args{10, 0}},
		{"10_1", args{10, 1}},
		{"10_2", args{10, 2}},
		{"10_8", args{10, 8}},
		{"10_16", args{10, 16}},
		{"10_32", args{10, 32}},
		{"10_64", args{10, 64}},
		{"10_128", args{10, 128}},
	}

	for _, tt := range tests {
		n, err := strconv.ParseUint(groupID, tt.args.base, tt.args.bizSize)
		if err != nil {
			t.Errorf("case: %s, got err: %s", tt.name, err)
		}
		t.Logf("case: %s, got n=>%+v", tt.name, n)
	}

	//
}
