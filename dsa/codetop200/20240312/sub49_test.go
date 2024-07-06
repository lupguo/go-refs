package _0240312

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	type args struct {
		strs []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{"t1", args{[]string{"a"}}, [][]string{{"a"}}},
		{"t2", args{[]string{"ab", "ba"}}, [][]string{{"ab", "ba"}}},
		{"t3", args{[]string{"ab", "aba"}}, [][]string{{"ab"}, {"aba"}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := groupAnagrams(tt.args.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %v, want %v", got, tt.want)
			}
		})
	}
}
