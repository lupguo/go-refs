package osx

import (
	"os"
	"testing"
)

func TestDelOpenFile(t *testing.T) {
	tempFile, err := os.CreateTemp("", "to_del")
	if err != nil {
		t.Errorf("os.CreateTemp got err: %s", err)
	}
	defer tempFile.Close()

	// 打开文件
	f, err := os.Open(tempFile.Name())
	if err != nil {
		t.Errorf("os.Open got err: %s", err)
	}
	t.Logf("f name: %s", f.Name())

	// 删除文件
	err = os.Remove(tempFile.Name())
	if err != nil {
		t.Errorf("os.Remove got err: %s", err)
	}
}
