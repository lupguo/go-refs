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

	t.Logf("new(Student) => %#v, %[1]p", s1)
	t.Logf("&Student{} => %#v, %[1]p", s2)
	t.Logf("s1==s2? => %t", s1 == s2)
}

func TestASM(t *testing.T) {
	// const (
	// 	Enone  = 0
	// 	Eio    = 1
	// 	Einval = 2
	// )
	// a := [...]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	// s := []string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
	// m := map[int]string{Enone: "no error", Eio: "Eio", Einval: "invalid argument"}
}
