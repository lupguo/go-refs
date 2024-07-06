package rangex

import (
	"testing"
)

func TestForRange(t *testing.T) {
	v1 := []int{1, 2, 3}
	v2 := map[int]bool{1: true, 2: false}
	v3 := make(chan int, 3)
	v3 <- 1
	v3 <- 2
	v3 <- 3
	close(v3)
	v4 := "hello中国"

	// slice
	for idx, val := range v1 {
		t.Logf("idx=%v, val=%v", idx, val)
	}

	// map
	for key, val := range v2 {
		t.Logf("key=%v, val=%v", key, val)
	}

	// chan
	for v := range v3 {
		t.Logf("recive from chan: %v", v)
	}

	// string
	for i, c := range v4 {
		t.Logf("i=%d, s=%c", i, c)
	}
}
