package main

import (
	"bufio"
	"fmt"
	"net"

	log "github.com/sirupsen/logrus"
	"x-learn/adv/nettcpx/bufiod"
)

// tcp client
// dial -> conn
// for scan { conn <- input}
func main() {
	conn, err := net.Dial("tcp", bufiod.BufioAddr)
	if err != nil {
		log.Fatalf("net dial got err: %v", err)
	}
	wr := bufio.NewWriter(conn)
	rd := bufio.NewReader(conn)

	for {
		// read from cmdline
		var str string
		_, err := fmt.Scanln(&str)
		if err != nil {
			log.Errorf("sscanln got err: %v", err)
			continue
		}
		log.Infof("client say: %v", str)

		// write to conn
		_, err = wr.WriteString(str)
		if err != nil {
			log.Errorf("cli write got err: %v", err)
			return
		}
		wr.WriteByte('\n')
		wr.Flush()

		// read
		rsp, err := rd.ReadString('\n')
		log.Infof("buff data: %v", rd.Buffered())
		log.Infof("client log rsp: %v", rsp)
	}
}
