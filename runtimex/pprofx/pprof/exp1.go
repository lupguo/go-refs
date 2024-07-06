package pprof

//
// func escapeExample1() func() int {
// 	x := 10 // x 在函数内部被分配
// 	return func() int {
// 		return x // 闭包中引用了外部变量 x，导致 x 逃逸到堆上
// 	}
// }
//
// // 赋值给全局变量
// var globalMap map[string]*int
//
// func escapeExample2() {
// 	x := 10               // x 在函数内部被分配
// 	globalMap["key"] = &x // x 的指针被存储在全局映射中，逃逸到堆上
// }
//
// // 返回内部变量指针地址
// func escapeExample3() *int {
// 	x := 10   // x 在函数内部被分配
// 	return &x // x 在函数外部被引用，逃逸到堆上
// }

// Bird 接口
type Bird interface {
	Fly()
}

type Duck struct {
	Name string
}

func (d *Duck) Fly() {
}

func main() {
	var b Bird
	duck := &Duck{}
	b = duck
	b.Fly()
}
