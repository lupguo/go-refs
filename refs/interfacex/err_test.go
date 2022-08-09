package interfacex

import (
	"log"
	"testing"
)

type MyError interface {
	Error()
}

type KError struct {
	Name string
}

func (K KError) Error() {
	log.Println("kError log...")
}

func TestErrPrint(t *testing.T) {
	var e, e1 MyError
	var err1 *KError
	var err2 KError
	t.Logf("%#v, nil=%t, e==e1?=>%t, &e=(%x)", e, e == nil, e == e1, &e) // 申明了，MyError接口，两个空接口也是相等的
	t.Logf("%#v, nil=%t", err1, err1 == nil)                             // 申明存储KError的指针变量
	t.Logf("%#v, nil=%t", err2, &err2 == nil)                            // 已申明，有内存地址

	// e2 := e.(*KError)
	// t.Logf("e2.Name:%s", e2.Name) // panic，因为e2是*KError的空指针，所以会报panic

	// e3 := e.(KError) // panic，因为e是接口，接口断言只能是指针类型
	// t.Logf("e3.Name:%s", e3.Name)

	e4 := &err2
	t.Logf("e4.Name:%s", e4.Name) // e4是有err2结构体对应的指针，err2有具体结构体的内存存储，所以e4.Name返回空

	e5 := new(KError)
	t.Logf("e5.Name:%s", e5.Name) // e5是初始化话的KError结构体对应的指针，有具体结构体的内存存储，所以e5.Name返回空
}

func NilAssert(e MyError) {
	e1 := e.(*KError)
	if e1 != nil {

	}
}
