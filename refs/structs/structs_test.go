package structs

import (
	"testing"
)

type Person struct {
	name string
	age  int
}

func TestStruct(t *testing.T) {
	var p1 Person
	p2 := new(Person)
	p3 := &Person{}
	p4 := Person{
		name: "",
		age:  0,
	}

	t.Logf("p1 %v, %s, pointer:%p\n", p1, p1.name, &p1)
	t.Logf("p2 %v, %s, pointer:%p\n", p1, p2.name, p2)
	t.Logf("p3 %v, %s, pointer:%p\n", p1, p3.name, p3)
	t.Logf("p4 %v, %s, pointer:%p\n", p1, p4.name, &p4)

	t.Logf("equal:\n")
	t.Logf("%t", &p1 == nil)
	t.Logf("%t", &p2 == nil)
}
