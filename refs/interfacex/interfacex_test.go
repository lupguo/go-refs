package interfacex

import (
	"io"
	"os"
	"testing"
)

type Fooer interface {
	Foo()
	ImplementsFooer()
}

type Bar struct{}

func (b Bar) ImplementsFooer() {}
func (b Bar) Foo()             {}

type Equaler interface {
	Equal(Equaler) bool
}

type T int

func (t T) Equal(u T) bool { return t == u } // does not satisfy Equaler
func TestT1Equal(t *testing.T) {
	a := T(2)
	b := 2
	t.Logf("%+v", a.Equal(T(b)))
}

type T2 int

func (t T2) Equal(u Equaler) bool { return t == u.(T2) } // satisfies Equaler
func TestT2Equal(t *testing.T) {
	a := T2(2)
	b := 2
	t.Logf("%+v", a.Equal(T2(b)))
}

type Opener interface {
	Open() io.Reader
}

type T3 struct{}

func (t T3) Open() *os.File {
	return nil
}

func TestConverSliceAB(t *testing.T) {
	type A []int
	type B []int

	a := A{1, 2}
	b := B(a)
	c := []int(a)

	for _, i := range b {
		t.Logf("%d", i)
	}
	for _, i := range c {
		t.Logf("%d", i)
	}
}

func TestConverSliceCD(t *testing.T) {
	type A int
	type B int
	type C []A
	type D []B

	x := C{1, 2}
	// y := D(x)
	_ = x
}

type Copyable interface {
	Copy() interface{}
}

type Value struct {
}

func (v Value) Copy() Value {
	return Value{}
}
