package synclib

import (
	"sync"
	"testing"
)

var aa string
var once sync.Once

func setup() {
	println("set aa")
	aa = "hello, world"
}

func doprint() {
	once.Do(setup)
	println(aa)
}

func TestOnceManyTimes(t *testing.T) {
	go doprint() // 执行setup
	go doprint() // 不会再执行setup函数，同时会阻塞等待第一个setup执行结束返回时候执行print(a)代码
}
