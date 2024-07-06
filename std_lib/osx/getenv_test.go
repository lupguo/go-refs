package osx

import (
	"fmt"
	"os"
	"testing"
)

func init() {
	if os.Getenv("OPEN_SARAMA_DEBUG_LOG") == "1" {
		fmt.Printf("[Good] OPEN_SARAMA_DEBUG_LOG is open....\n")
	}
}

func TestGetEnv(t *testing.T) {
	if os.Getenv("OPEN_SARAMA_DEBUG_LOG") == "1" {
		t.Logf("bingo get OPEN_SARAMA_DEBUG_LOG success")
		return
	}

	t.Logf("get GOPATH=>%s", os.Getenv("GOPATH"))

	t.Logf("CLICOLOR==1 ,got %t", os.Getenv("CLICOLOR") == "1")

	// for _, envstr := range os.Environ() {
	// 	t.Logf("evnStr=>%s", envstr)
	// }
}
