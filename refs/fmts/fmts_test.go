package fmts

import (
	"encoding/json"
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSscanf(t *testing.T) {
	var id uint64
	sscanf, err := fmt.Sscanf("id:100", `id:%d`, &id)
	assert.Nil(t, err)
	t.Logf("sscanf=%d, id:%d", sscanf, id)
}

func Test520(t *testing.T) {
	s := "Lover"
	for i, c := range s {
		fmt.Printf("i=>%d, n=>%d, c=>%[2]c\n", i, c)
	}
	fmt.Println('L' + 'o' + 'v' + 'e' + 'r')
}

func TestFormt(t *testing.T) {
	t.Logf("%0.2f", 0.0)
	t.Logf("%0.2f", complex(0, 3))
	ia := 15
	t.Logf("%0.2f", float64(ia))
	ib := 10
	t.Logf("%0.2d", ib)
}

func TestPrintV(t *testing.T) {
	type ttype struct {
		A int
		B string
	}

	v := ttype{100, "hello"}

	vs := []*ttype{
		{1, "aaa"},
		{2, "bbb"},
	}

	t.Logf("struct: %v, %+[1]v, %#[1]v", v)
	t.Logf("pointer: %v, %+[1]v, %#[1]v", &v)

	// slice
	t.Logf("raw: %v, %+[1]v, %#[1]v", vs)
	t.Logf("slice: %v, %+[1]v, %#[1]v", &vs)

	// json print
	b, err := json.Marshal(vs)
	if err != nil {
		t.Error(err)
	}
	t.Logf("%s", b)

}
