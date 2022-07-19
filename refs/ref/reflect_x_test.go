package ref

import (
	"reflect"
	"testing"
)

func TestReflectValueAndType(t *testing.T) {
	// defer func() {
	// 	r := recover()
	//
	// }()
	x := 3.14
	t.Logf("type=>%+v, val=>%+v", reflect.TypeOf(x), reflect.ValueOf(x))
	t.Logf("string => %s", reflect.ValueOf(x).String())

	// v3.14
	v := reflect.ValueOf(x)
	t.Log("type:", v.Type())
	t.Log("kind is float64:", v.Kind() == reflect.Float64)
	t.Log("value float64:", v.Float())

	// range slice
	y := []uint64{1, 2, 3}
	k := reflect.ValueOf(y)
	t.Log("kind is slice:", k.Kind() == reflect.Slice)
	t.Log("kind type is :", k.Type())
	t.Logf("interface is %+v, %[1]t", k.Interface())
	t.Log("value slice 1 is :", k)
	t.Log("value slice 2 is :", k.Slice(0, k.Len()-1))

	// ref val can set?
	var z float64 = 3.4
	p := reflect.ValueOf(&z) // Note: take the address of x.
	t.Log("type of p:", p.Type())
	t.Log("settability of p:", p.CanSet())
	t.Log("settability of p.Elem():", p.Elem().CanSet())

	// ref.type -> interface
	var w = []uint64{1, 2, 3}
	rvw := reflect.ValueOf(w)
	wi := reflect.ValueOf(w).Interface()
	for i, iv := range wi.([]uint64) {
		t.Logf("slice[%d] => %d ", i, iv)
	}

	for i:=0; i< rvw.Len(); i++ {
		t.Logf("slice range[%d] => %v", i, rvw)
	}

	// ks := k.Slice(0, k.Len())
	// for _, kvv := range ks {
	// 	t.Logf("kvv => %v", kvv)
	// }
	//
	// for k, v := range []uint64(k.Type()) {
	//
	// }
}
