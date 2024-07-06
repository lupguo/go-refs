package reflectt

import (
	"fmt"
	"reflect"
	"testing"
)

func TestCopyByPointer(t *testing.T) {
	v := "hello"
	i := 100
	o := &User{Name: "clark"}

	// before
	t.Logf("before v=>%s, i=>%d, o=>%v", v, i, o)
	UpdateByReflect(&v, "world")
	// UpdateByReflect(&i, 88)
	// UpdateByReflect(o, User{"terry"})
	// after
	t.Logf("after v=>%s, i=>%d, o=>%v", v, i, o)
}

// 弊端，无法穷举
func UpdateBySwitchType(v interface{}) {
	switch p := v.(type) {
	case *int:
		*p = 42
	case *string:
		*p = "world"
	case *struct{ Name string }:
		*p = struct{ Name string }{Name: "terry"}
	default:
		panic("v can't be updated")
	}
}

func UpdateByReflect(v interface{}, val string) {
	d := reflect.ValueOf(v).Elem()
	fmt.Printf("d=>%v\n", d)
	fmt.Printf("d.Addr()=>%v\n", d.Addr())
	if d.CanSet() {
		d.Set(reflect.ValueOf(val))
	}
}
