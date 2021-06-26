package forpointers

import (
	"testing"
)

type User struct {
	ID int
}

func TestName(t *testing.T) {
	us := []*User{
		{1},
		{2},
		{3},
	}

	m := make(map[int]*User)
	for _, u := range us {
		t.Logf("Pu=%p, Pv=>%p, u=>%+v",&u, u, *u)
		m[u.ID] = u
	}

	for k, u := range m {
		t.Logf("k=>%d,u=>%d", k, u)
	}

	// ---
	us2 := []User{
		{1},
		{2},
		{3},
	}
	m2 := make(map[int]User)
	for _, u := range us2 {
		t.Logf("Po=>%p, u=>%+v", &u, u)
		m2[u.ID] = u
	}
	for k, u := range m2 {
		t.Logf("k=>%d,u=>%d", k, u)
	}
}
