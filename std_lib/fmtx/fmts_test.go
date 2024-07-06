package fmtx

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/hold7techs/goval"
	"github.com/lupguo/go-render/render"
)

type User struct {
	id   int
	Name string
}

// 结构体打印
// fmts_test.go:47: struct: {1 user1}, {id:1 Name:user1}, fmts.User{id:1, Name:"user1"}
// fmts_test.go:48: pointer: &{2 user2}, &{id:2 Name:user2}, &fmts.User{id:2, Name:"user2"}
func TestPrintV(t *testing.T) {
	u1 := User{1, "user1"}
	u2 := &User{2, "user2"}

	// %v 不带属性名称 {1 user1}
	// %+v 带属性名称，{id:1 Name:user1}
	// %#v 会带上Go语言结构体名称 fmts.User{id:1, Name:"user1"}
	t.Logf("struct: %v, %+[1]v, %#[1]v", u1)
	t.Logf("pointer: %v, %+[1]v, %#[1]v", u2)
}

// 指针结构体切片打印
// fmts_test.go:66: raw: [0x140001302a0 0x140001302b8], [0x140001302a0 0x140001302b8], []*fmts.User{(*fmts.User)(0x140001302a0), (*fmts.User)(0x140001302b8)}
// fmts_test.go:67: slice: &[0x140001302a0 0x140001302b8], &[0x140001302a0 0x140001302b8], &[]*fmts.User{(*fmts.User)(0x140001302a0), (*fmts.User)(0x140001302b8)}
// fmts_test.go:74: [{"Name":"user1"},{"Name":"user2"}] // Json序列化因为Name是导出的，所以支持打印出来
func TestPrintSlice(t *testing.T) {
	// 如果是结构体指针切片，%v 都只会打印切片元素的值而不会打印指针的结果，
	users := []*User{
		{1, "user1"},
		{2, "user2"},
	}

	// slice - %v和%+v没有什么区别，%#v会打印内部数据结构
	t.Logf("raw: %v, %+[1]v, %#[1]v", users)
	t.Logf("slice: %v, %+[1]v, %#[1]v", &users)

	// json marshal print
	b, err := json.Marshal(users)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", b)
}

type Member struct {
	id   int
	name string
}

// 让Member结构体实现Strings接口，支持%v打印
func (m *Member) String() string {
	return fmt.Sprintf(`{id:%d,Name:%s}`, m.id, m.name)
}

// 通过String()输出有个好处，就是当结构体为nil时候会正常打印nil而不会直接panic
func TestPrintNilMember(t *testing.T) {
	var m *Member
	// t.Logf("%v", m.id) // will panic
	t.Logf("%v", m)
}

// fmts_test.go:92: raw: [{id:1,Name:MemberA} {id:2,Name:MemberB}], [{id:1,Name:MemberA} {id:2,Name:MemberB}], []*fmts.Member{(*fmts.Member)(0x1400000e2a0), (*fmts.Member)(0x1400000e2b8)}
// 可以看到，
//
//	%v、%+v -> 按String()方法打印
//	%#v -> 按Go结构体类型打印，所以内部还是指针
//	结论: 日志打印，通常用%v或%+v就可以，当需要附带属性名称时候，使用%+v；%#v通常用于查看Go原始的值，比如查看指针的值
func TestPrintSlice2(t *testing.T) {
	members := []*Member{
		{1, "MemberA"},
		{2, "MemberB"},
	}
	t.Logf("raw: %v, %+[1]v, %#[1]v", members)
}

type Extend struct {
	School string
	Home   string
}

// 结构体指针嵌套
type NestUser struct {
	ID     int
	Name   string
	Favs   []int
	Extend *Extend
	Family map[string]string
}

// fmts_test.go:106: %v => &{100 UserA [4 5] 0x1400014c040 map[father:fa mother:mo]}
// fmts_test.go:107: %+v => &{ID:100 Name:UserA Favs:[4 5] Extend:0x1400014c040 Family:map[father:fa mother:mo]}
// fmts_test.go:108: %#v => &fmts.NestUser{ID:100, Name:"UserA", Favs:[]int{4, 5}, Extend:(*fmts.Extend)(0x1400014c040), Family:map[string]string{"father":"fa", "mother":"mo"}}
// 从这个结果看得出很明显的效果，Go不会迭代打印，通常情况日志用%+v包含结构体属性标签更适合Debug
func TestPrintNestStruct(t *testing.T) {
	u1 := &NestUser{
		ID:   100,
		Name: "UserA",
		Favs: []int{4, 5},
		Extend: &Extend{
			School: "SchoolXiao",
			Home:   "HomeTang",
		},
		Family: map[string]string{
			"mother": "mo",
			"father": "fa",
		},
	}
	t.Logf(`%%v => %v`, u1)
	t.Logf(`%%+v => %+v`, u1)
	t.Logf(`%%#v => %#v`, u1)

	t.Logf("goVal=>%s", goval.ToTypeString(u1))
}

// 尝试用反射
// 通过反射获取到v的值表示形式，会不断迭代下去
// fmts_test.go:136: netUser=>&fmts.NestUser{ID: 100, Name: "UserA", Favs: []int{4, 5}, Extend: &fmts.Extend{School: "SchoolXiao", Home: "HomeTang"}, Family: map[string]string{"mother": "mo", "father": "fa"}}
// fmts_test.go:148: netMap=>map[string]*fmts.Extend{"ext2": &fmts.Extend{School: "School2", Home: "Home2"}, "ext1": &fmts.Extend{School: "School1", Home: "Home1"}}
func TestReflectGovalPrint(t *testing.T) {
	u1 := &NestUser{
		ID:   100,
		Name: "UserA",
		Favs: []int{4, 5},
		Extend: &Extend{
			School: "SchoolXiao",
			Home:   "HomeTang",
		},
		Family: map[string]string{
			"mother": "mo",
			"father": "fa",
		},
	}

	// 嵌套打印
	t.Logf("netUser=>%s", goval.ToTypeString(u1))

	m := map[string]*Extend{
		"ext1": {
			School: "School1",
			Home:   "Home1",
		},
		"ext2": {
			School: "School2",
			Home:   "Home2",
		},
	}
	t.Logf("netMap=>%s", goval.ToTypeString(m))
}

// fmts_test.go:173: netUser=>(*fmts.NestUser){ID:100, Name:"UserA", Favs:[]int{4, 5}, Extend:(*fmts.Extend){School:"SchoolXiao", Home:"HomeTang"}, Family:map[string]string{"father":"fa", "mother":"mo"}}
// fmts_test.go:185: netMap=>map[string]*fmts.Extend{"ext1":(*fmts.Extend){School:"School1", Home:"Home1"}, "ext2":(*fmts.Extend){School:"School2", Home:"Home2"}}
func TestReflectRenderPrint(t *testing.T) {
	u1 := &NestUser{
		ID:   100,
		Name: "UserA",
		Favs: []int{4, 5},
		Extend: &Extend{
			School: "SchoolXiao",
			Home:   "HomeTang",
		},
		Family: map[string]string{
			"mother": "mo",
			"father": "fa",
		},
	}

	// 嵌套打印
	t.Logf("netUser=>%s", render.Render(u1))

	m := map[string]*Extend{
		"ext1": {
			School: "School1",
			Home:   "Home1",
		},
		"ext2": {
			School: "School2",
			Home:   "Home2",
		},
	}
	t.Logf("netMap=>%s", render.Render(m))
}

func BenchmarkRender(b *testing.B) {
	u1 := &NestUser{
		ID:   100,
		Name: "UserA",
		Favs: []int{4, 5},
		Extend: &Extend{
			School: "SchoolXiao",
			Home:   "HomeTang",
		},
		Family: map[string]string{
			"mother": "mo",
			"father": "fa",
		},
	}
	for i := 0; i < b.N; i++ {
		render.Render(u1)
	}
}

func BenchmarkGovalToString(b *testing.B) {
	u1 := &NestUser{
		ID:   100,
		Name: "UserA",
		Favs: []int{4, 5},
		Extend: &Extend{
			School: "SchoolXiao",
			Home:   "HomeTang",
		},
		Family: map[string]string{
			"mother": "mo",
			"father": "fa",
		},
	}
	for i := 0; i < b.N; i++ {
		goval.ToString(u1)
	}
}
