package gomaxprocs

import (
	"fmt"
	"log"
	"runtime"
	"testing"
	"time"
)

func TestGomaxprocs(t *testing.T) {
	runtime.GOMAXPROCS(2)

	go func() {
		if v := <-time.Tick(time.Second); &v != nil {
			log.Fatal("time over")
		}
	}()

	for {
		go fmt.Print(0)
		fmt.Print(1)
	}
}
