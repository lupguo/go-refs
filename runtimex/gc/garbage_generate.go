package main

func main() {
	NewDuck()
}

type Duck struct {
}

//go:noinline
func NewDuck() *Duck {
	return &Duck{}
}
