package maps

import (
	"sort"
	"testing"
)

func TestRangeSortedMap(t *testing.T) {
	v := map[int]int{
		1: 1,
		2: 2,
		3: 3,
	}
	for k, vv := range v {
		t.Logf("%d=>%v", k, vv)
	}
	for k, vv := range sort.Sort(v) {
		t.Logf("%d=>%v", k, vv)
	}
}
