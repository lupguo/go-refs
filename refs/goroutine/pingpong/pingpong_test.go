package pingpong

import (
	"testing"
	"time"
)

func TestPingPongForever(t *testing.T) {
	ch1, ch2 := make(chan string), make(chan string)

	go func() {
		for v := range ch2 {
			_ = v
			// fmt.Printf("%s!\n", v)
			ch1 <- "ping"
		}
	}()

	cnt := 0
	go func() {
		for v := range ch1 {
			cnt++
			_ = v
			// fmt.Printf("%s...\n", v)
			ch2 <- "pong"
		}
	}()

	go func() {
		ch1 <- "ping"
		for {
			select {
			case <-time.Tick(time.Second):
				t.Logf("%s, cnt=>%d", time.Now(), cnt)
				cnt = 0
			}
		}
	}()

	select {}
}
