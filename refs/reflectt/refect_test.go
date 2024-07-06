package reflectt

import (
	"reflect"
	"testing"
)

type User struct {
	Name string
	Age  uint32
}

func TestRefectType(t *testing.T) {
	u := User{"clark", 100}
	rt := reflect.TypeOf(u)
	t.Logf("rt.String()=>%s, , rt.Name()=>%s", rt.String(), rt.Name())

	var fields []string
	for i := 0; i < rt.NumField(); i++ {
		fields = append(fields, rt.Field(i).Name)
	}
	t.Log(fields)
	t.Log(rt.FieldByIndex([]int{0}))
}

func TestPointerRefectType(t *testing.T) {
	u := &User{"clark", 100}
	rt := reflect.TypeOf(*u)
	t.Logf("rt.String()=>%v, , rt.Name()=>%v", rt.String(), rt.Name())

	var fields []string
	for i := 0; i < rt.NumField(); i++ {
		fields = append(fields, rt.Field(i).Name)
	}
	t.Log(fields)
}
