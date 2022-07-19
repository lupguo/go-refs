package os

import (
	"os/exec"
	"testing"
)

func TestLs(t *testing.T) {
	cmd := exec.Command("ls", "/tmp")
	err := cmd.Run()
	if err != nil {
		t.Errorf("run got err: %s",err)
	}
}



