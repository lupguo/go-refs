package interfacex

import (
	"testing"
)

func TestAssert(t *testing.T) {
	var a interface{}
	a = 100
	v, ok := a.(uint64)
	t.Logf("v=%+v, ok=%+v", v, ok)
	v1, ok := a.(int)
	t.Logf("v=%+v, ok=%+v", v1, ok)
	v2, ok := a.(float64)
	t.Logf("v=%+v, ok=%+v", v2, ok)

	// if reflect.TypeOf(a).Kind() == reflect.Uint64 {
	// 	reflect.ValueOf(a).SetUint()
	// }
	//
	// // 强转
	// var b interface{}

}
