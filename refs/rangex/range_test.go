package rangex

import "testing"

func TestRange(t *testing.T) {
	a := 1
	b := 2
	c := 3
	ss := []*int{&a, &b, &c}
	sm := make(map[int]int)
	for i, s := range ss {
		t.Logf("s=>%d, pinter s=>%p", s, &s)
		sm[i] = *s
	}
	t.Logf("sm => %#v, len => %d", sm, len(sm))
}

func TestRangeSlice(t *testing.T) {
	type accType struct {
		id   int
		name string
	}
	m := make(map[int]accType)
	accs := []accType{{1, "zhang"}, {2, "wang"}, {3, "li"}}
	for _, acc := range accs {
		if acc.id == 2 {
			acc.name = "mock"
		}
		m[acc.id] = acc
	}
	t.Logf("%+v", m)
}

func TestRangeNil(t *testing.T) {
	var s1 []int
	for i, v := range s1 {
		t.Log(i, v)
	}
	var m1 map[int]int
	for i, i2 := range m1 {
		t.Log(i, i2)
	}
	// --
	var m2 = make(map[int]int)
	for i, i2 := range m2 {
		t.Log(i, i2)
	}
}
