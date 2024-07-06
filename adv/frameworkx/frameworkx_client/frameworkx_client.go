package main

import (
	"bufio"
	"net"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
	"golang.org/x/sync/errgroup"
)

func main() {

	//
	egp := errgroup.Group{}
	egp.SetLimit(5)
	for i := 0; i < 10; i++ {
		egp.Go(func() error {
			err := connectService("127.0.0.1:8033")
			if err != nil {
				log.Errorf("connect service : %v", err)
			}
			return nil
		})
	}

	if err := egp.Wait(); err != nil {
		log.Printf("egp wait err: %v", err)
	}

	log.Println("client request fin")
}

func connectService(ipaddr string) error {
	conn, err := net.Dial("tcp", ipaddr)
	if err != nil {
		return errors.Wrapf(err, "dial tcp got er")
	}
	tcpConn := conn.(*net.TCPConn)
	err = clientRequest(tcpConn)
	if err != nil {
		return errors.Wrapf(err, "request got err")
	}

	log.Println("connectService done")
	return nil
}

func clientRequest(conn *net.TCPConn) error {

	reader := bufio.NewReader(conn)
	data, err := reader.ReadString('\n')
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	if err != nil {
		return errors.Wrapf(err, "conn read got err")
	}

	log.Printf("client request done: %s", data)

	return nil
}
