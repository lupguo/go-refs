package main

import (
	"fmt"
	"net"
)

func handleConnection(conn net.Conn) {
	// Handle RPC request
}

func main() {
	listener, err := net.Listen("tcp", "localhost:8033")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("Server listening on localhost:8080")

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			continue
		}
		go handleConnection(conn) // Handle connection asynchronously
	}

}
