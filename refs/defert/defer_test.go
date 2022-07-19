package defert

import (
	"fmt"
	"testing"
)

func TestDefer(t *testing.T) {
	a := 1
	defer func() {
		fmt.Printf("a1=%d", a)
	}()
	defer func(a int) {
		fmt.Printf("a2=%d", a)
	}(a)
	defer func(a *int) {
		fmt.Printf("a3=%d", *a)
	}(&a)
	if true {
		defer fmt.Println(1)
	}
	fmt.Println("func p1")
	defer fmt.Println(2)
	defer fmt.Println(3)
	a++
}
