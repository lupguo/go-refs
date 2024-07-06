package main

import (
	"fmt"
	"net"
)

func main() {
	// 创建TCP监听器
	listener, err := net.Listen("tcp", "localhost:8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Listening on localhost:8080")

	// 接受连接并处理
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}
		go handleConnection(conn)
	}
}

func handleConnection(conn net.Conn) {
	// 处理连接
	defer conn.Close()

	// 在全连接队列中处理连接
	// ...
}
