package osx

import (
	"os"
	"testing"
)

func TestMkdir(t *testing.T) {
	dirname := "testdir"
	// err := os.MkdirAll(dirname, 0750)
	// if err != nil {
	// 	t.Errorf("3 err: %s",err)
	// }

	// err = os.Mkdir(dirname, 0750)
	// if err != nil {
	// 	t.Errorf("1 err: %s",err)
	// }

	//
	// err = os.Mkdir(dirname, 0750)
	// if err != nil {
	// 	t.Errorf("2 err: %s",err)
	// }

	if err := os.MkdirAll(dirname, 0750); err != nil {
		t.Errorf("3 err: %s", err)
	}
}
