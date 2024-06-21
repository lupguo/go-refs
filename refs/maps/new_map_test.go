package maps

import (
	"testing"
)

type MetaData map[string][]byte

func TestPanic(t *testing.T) {
	msg := &MetaData{}
	(*msg)["hello"] = []byte("hello")
	t.Logf("%+v", msg)

}
