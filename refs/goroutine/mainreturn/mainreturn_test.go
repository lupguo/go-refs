package mainreturn

import (
	"testing"
	"time"
)

func TestMainReturn(t *testing.T) {
	go func() {
		// 主Gfn
		go func() {
			// 子G fn
			for {
				t.Log(time.Now())
				time.Sleep(time.Second)
			}
		}()

		// 主G退出，子G也没有了?
		time.Sleep(3 * time.Second)
		t.Log("Out goroutine return...")
		return
	}()

	select {}
}
