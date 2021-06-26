package os

import (
	"fmt"
	"os/exec"
	"testing"
)

func TestLs(t *testing.T) {
	cmd := exec.Command("ls", "")
	message := "hello"
	fmt.Println(message)
}



