package sort

import (
	"fmt"
	"sort"
	"testing"
)

func TestSort(t *testing.T) {
	people := []struct {
		Name string
		Age  int
	}{
		{"Gopher", 7},
		{"Alice", 55},
		{"Vera", 24},
		{"Bob", 75},
	}
	sort.Slice(people, func(i, j int) bool { return people[i].Name < people[j].Name })
	fmt.Println("By name:", people)
	sort.Slice(people, func(i, j int) bool { return people[i].Age < people[j].Age })
	fmt.Println("By age:", people)
}

type user struct {
	name string
	age  int
}
type kv struct {
	key  int
	user user
}

type ms []kv

func (m ms) Len() int {
	return len(m)
}

func (m ms) Less(i, j int) bool {
	return m[i].user.age < m[j].user.age
}

func (m ms) Swap(i, j int) {
	m[i], m[j] = m[j], m[i]
}

func TestSortMap(t *testing.T) {
	m := map[int]user{
		2: {"u2", 12},
		1: {"u1", 11},
		4: {"u4", 14},
		5: {"u5", 15},
		3: {"u3", 13},
	}

	s := make(ms, len(m))

	i := 0
	for k, u := range m {
		s[i] = kv{k, u}
		i++
	}

	sort.Sort(s)

	for _, u := range s {
		t.Logf("map[%d]=>%v\n", u.key, u.user)
	}

	// for k := 0; k < 100; k++ {
	// 	i := 1
	// 	for id := 1; id <= len(m); id++ {
	// 		if i != id {
	// 			t.Fatalf("id=%d, id#%d, user=>%v", i, id, m[id])
	// 		}
	// 		i++
	// 	}
	// }
}
