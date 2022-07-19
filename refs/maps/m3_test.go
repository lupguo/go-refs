package maps

import (
	"fmt"
	"testing"
)

func TestPrintMap(t *testing.T) {
	type st struct {
		ID   uint64
		name string
	}

	var mmap = map[string]*st{
		"k1": {100, "zhangsan"},
		"k2": {101, "lisi"},
	}

	fmt.Printf("%+v\n", mmap)
	fmt.Printf("%#v\n", mmap)
	fmt.Printf("%v\n", mmap)
}

func TestMap4(t *testing.T) {
	type user struct {
		ID uint64
	}

	fn := func() map[uint64]user {
		return nil
	}
	umap := fn()

	t.Log(umap[100].ID)
}

func TestMap5(t *testing.T) {
	type user struct {
		ID uint64
	}

	var umap map[uint64]user

	t.Logf("%+v", umap[100])

	umap = make(map[uint64]user)
	umap[10] = user{99}
	t.Logf("%+v", umap[99])
}
