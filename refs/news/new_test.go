package news

import "testing"

type Person struct {
	Name string
	Id   int
	Fav  []string
}

type Student struct {
	Person
	Class string
}

func TestNewAndVar(t *testing.T) {
	s1 := new(Student)
	s2 := &Student{}

	t.Logf("new(Student) => %#v", s1)
	t.Logf("&Student{} => %#v", s2)
}
