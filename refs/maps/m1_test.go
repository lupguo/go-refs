package maps

import (
	"fmt"
	"testing"
)

type cmdHandler func() error

func h1() error {
	fmt.Println("h1")
	return nil
}

func h2() error {
	fmt.Println("h2")
	return nil
}

func TestUndefined(t *testing.T) {
	cmdMap := make(map[string]cmdHandler)
	cmdMap["h1"] = h1
	cmdMap["h2"] = h2
}

func TestValExist(t *testing.T) {
	m := make(map[int]int)
	m[100] = 300
	k, ok := m[1]
	t.Log(k, ok)
	k, ok = m[100]
	t.Log(k, ok)

	t.Log(m[1])
}

func TestDeleteMap(t *testing.T) {
	m := make(map[int]int)
	delete(m , 1)
	delete(m , 2)
}