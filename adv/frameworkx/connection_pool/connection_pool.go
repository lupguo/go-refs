package main

import (
	"bytes"
	"fmt"
	"net"
	"net/http"
	_ "net/http/pprof"
	"time"

	log "github.com/sirupsen/logrus"
)

func handleConnection(pool chan net.Conn, conn net.Conn) {
	buf := bytes.NewBufferString(fmt.Sprintf("%s\n", time.Now().String()))
	_, err := conn.Write(buf.Bytes())
	if err != nil {
		return
	}
	if err != nil {
		return
	}
	// 模拟耗时
	time.Sleep(1 * time.Second)
	log.Println("server handle conn done")
	pool <- conn
}

func main() {
	// 注册pprof路由
	go func() {
		log.Println("pprof on http://localhost:6060/debug/pprof/")
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	listener, err := net.Listen("tcp", ":8033")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on localhost:8033")

	// Pre-create a pool of connections
	connectionPool := make(chan net.Conn, 10)
	for i := 0; i < 10; i++ {
		log.Printf("listen[%v]...", i)
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		connectionPool <- conn
	}

	for {
		conn := <-connectionPool
		log.Println("conn handle...")
		go handleConnection(connectionPool, conn) // Handle connection asynchronously

	}
}
