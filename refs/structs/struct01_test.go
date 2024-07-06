package structs

import (
	"testing"
)

type ABC struct {
	A struct {
		B struct {
			b1 string
			b2 int
		}
	}
}

type DEF struct {
	D struct {
		E *struct {
			b1 string
			b2 int
		}
	}
}

func TestDEF(t *testing.T) {
	def1 := &DEF{}
	t.Logf("%+v", def1)

	def2 := new(DEF)
	t.Logf("%+v", def2)

	t.Logf("case1 a.b.c b1=%v, b2=%v", def1.D.E.b2, def1.D.E.b2)
	t.Logf("case2 a.b.c b1=%v, b2=%v", def2.D.E.b2, def2.D.E.b1)
}

func TestABC(t *testing.T) {
	abc1 := &ABC{}
	t.Logf("%+v", abc1)

	abc2 := new(ABC)
	t.Logf("%+v", abc2)

	t.Logf("case1 a.b.c b1=%v, b2=%v", abc1.A.B.b1, abc1.A.B.b2)
	t.Logf("case2 a.b.c b1=%v, b2=%v", abc2.A.B.b1, abc2.A.B.b2)
}

type K struct {
	Num *int
}

func TestDefaultValue(t *testing.T) {
	k1 := new(K)
	n := 0
	k2 := &K{Num: &n}
	t.Logf("%#v", k1)
	t.Logf("%#v", k2)
}
