package clock

import (
	"io"
	"net"
	"sync"
	"testing"
	"time"
)

func TestClockCli(t *testing.T) {
	wg := sync.WaitGroup{}
	for i := 0; i < 10000; i++ {
		wg.Add(1)
		go func(id int) {
			connClockCli(t, id)
			wg.Done()
		}(i)
	}
	wg.Wait()
}

func connClockCli(t *testing.T, id int) {
	t.Logf("start client[%d]...\n", id)

	conn, err := net.DialTimeout("tcp", ":9095", 3*time.Second)
	if err != nil {
		t.Logf("conn got err: %s", err)
		return
	}
	defer conn.Close()

	t.Logf("client[%d] begin revice...\n", id)
	// for {
	// buf := make([]byte, 1024)
	// buffer := bytes.NewBuffer(buf)
	rb, err := io.ReadAll(conn)
	if err != nil {
		t.Errorf("client[%d] conn read got err: %s", id, err)
		return
	}
	t.Logf("client[%d] read data=\n%s", id, rb)
	// }
}
