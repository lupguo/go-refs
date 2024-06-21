package main

import (
	"context"
	"io"
	"net"
	"time"

	"github.com/pkg/errors"
	"x-learn/advance/klog/log"
)

func main() {
	address := "127.0.0.1:3351"
	listen, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatalf("listen got err: %s", err)
	}
	log.Infof("server listen on: %s", address)
	for {
		conn, err := listen.Accept()
		if err != nil {
			log.Errorf("listen got err: %s", err)
			return
		}
		go func(conn net.Conn) {
			defer conn.Close()

			// ctx
			background := context.Background()
			ctx, cancel := context.WithTimeout(background, 30*time.Second)
			defer cancel()

			// handle
			if err := Handle(ctx, conn); err != nil {
				log.Error(errors.Wrap(err, "handle got err"))
			}
		}(conn)
	}
}

func Handle(ctx context.Context, conn net.Conn) error {
	// go log
	log.Infof("begin receive data from conn...")
	defer log.Infof("conn finished.")

	for {
		select {
		case <-ctx.Done():
			return errors.Errorf("ctx got err, %v", ctx.Err())
		default:
			_, err := io.WriteString(conn, time.Now().Format("15:04:05.000\n"))
			if err != nil {
				return err
			}
			time.Sleep(time.Millisecond * 200)
		}
	}
}
