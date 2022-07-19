package clock

import (
	"io"
	"net"
	"testing"
	"time"
)

func TestClockSrv(t *testing.T) {
	// listen
	listener, err := net.Listen("tcp", ":9095")
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("listen: %s...", listener.Addr().String())

	// accept
	for {
		conn, err := listener.Accept()
		if err != nil {
			t.Logf("accept got err: %s", err)
			continue
		}

		// handle
		go handler(t, conn)
	}
}

func handler(t *testing.T, conn net.Conn) {
	defer conn.Close()

	// hand conn
	t.Logf("%s->%s\n", conn.RemoteAddr(), conn.LocalAddr())

	// over tick
	tick := time.Tick(time.Second)
	for {
		_, err := io.WriteString(conn, time.Now().Format("2006-01-02 15:04:05\n"))
		if err != nil {
			t.Logf("handle write string got err: %s", err)
			return
		}
		select {
		case <-tick:
			t.Logf("remote[%s] handle finished!", conn.RemoteAddr())
			return
		default:
			time.Sleep(time.Second)
		}
	}
}
