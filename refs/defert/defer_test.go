package defert

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	a := 1
	defer func(a *int) {
		fmt.Printf("a=%d\n", *a)
		fmt.Println(1)
	}(&a)
	defer fmt.Println(2)
	defer fmt.Println(3)
	a++
}
