package main

import (
	"fmt"

	_ "x-learn/refs/importx/pkga"
	_ "x-learn/refs/importx/pkgb"
)

func init() {
	fmt.Printf("main init(3)")
}

func main() {
	fmt.Printf("main func")
}
