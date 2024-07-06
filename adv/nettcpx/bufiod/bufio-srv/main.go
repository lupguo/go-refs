package main

import (
	"bufio"
	"net"

	log "github.com/sirupsen/logrus"
	"x-learn/adv/nettcpx/bufiod"
)

// listen+bind, for {conn := accept go handle(conn)}
func main() {
	// listen+bind
	listen, err := net.Listen("tcp", bufiod.BufioAddr)
	if err != nil {
		log.Fatalf("Listen got err: %v", err)
	}

	for {
		// accept
		conn, err := listen.Accept()
		if err != nil {
			log.Errorf("Accept got err: %v", err)
			continue
		}

		// handle
		go handleConn(conn)
	}

}

// read + write
func handleConn(conn net.Conn) {
	defer conn.Close()

	// read
	rd := bufio.NewReader(conn)
	wr := bufio.NewWriter(conn)

	// for {
	// read from conn
	got, err := rd.ReadString('\n')
	if err != nil {
		log.Errorf("read conn got err %v", err)
		return
	}
	log.Infof("srv recv: %v", got)

	// response
	wr.WriteString("service resp:")
	wr.WriteString(got)
	wr.Flush()
	// }

}
