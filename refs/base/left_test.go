package base

import (
	"fmt"
	"testing"
)

func TestLeft(t *testing.T) {
	for i := 0; i < 10; i++ {
		ii := 1 << i
		fmt.Printf("%b, %[1]T\n", ii)
	}
}
