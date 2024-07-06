package main

import (
	"io"
	"net"
	"os"
	"os/signal"

	"github.com/hold7techs/goval"
	"x-learn/advanced/klog/log"
)

func main() {
	// conn
	address := "127.0.0.1:3351"
	conn, err := net.Dial("tcp", address)
	if err != nil {
		log.Fatalf("dail got err: %s", err)
	}
	log.Infof("conn=>%s", goval.ToTypeString(conn))
	defer conn.Close()

	// exit
	c := make(chan os.Signal, 1)
	go func(ch chan<- os.Signal) {
		signal.Notify(ch, os.Interrupt)
	}(c)

	// conn close
	done := make(chan struct{})
	go func() {
		select {
		case <-c:
			conn = conn.(*net.TCPConn)
			conn.Close()
			done <- struct{}{}
		}
	}()

	// read
	buf := make([]byte, 3)
	n, err := io.CopyBuffer(os.Stdout, conn, buf)
	if err != nil {
		log.Errorf("client read from conn got err: %s", err)
		return
	}
	log.Infof("netcat got: %d", n)

	<-done
}
