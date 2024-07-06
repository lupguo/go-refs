package main

import (
	"fmt"
	"net"
	"time"
)

func handleConnection(hconn net.Conn) {
	conn := hconn.(*net.TCPConn)
	defer conn.Close()

	// 设置保活参数
	err := conn.SetKeepAlive(true)
	if err != nil {
		fmt.Println("Error setting keep alive:", err)
		return
	}

	// 设置保活间隔和最大重试次数（可根据需求进行调整）
	err = conn.SetKeepAlivePeriod(5 * time.Minute)
	if err != nil {
		fmt.Println("Error setting keep alive period:", err)
		return
	}

	// 处理连接逻辑
	// 在这里可以根据具体需求进行数据的读取和写入操作
	// 示例代码中只是简单地读取并打印客户端发送的数据
	buffer := make([]byte, 1024)
	n, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading data:", err)
		return
	}
	fmt.Println("Received data:", string(buffer[:n]))
}

func main() {
	// 监听指定端口
	listener, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Println("Error listening:", err)
		return
	}
	defer listener.Close()

	fmt.Println("TCP server is listening on port 8080")

	for {
		// 接受客户端连接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("Error accepting connection:", err)
			return
		}

		fmt.Println("New client connected:", conn.RemoteAddr())

		// 启动新的goroutine处理连接
		go handleConnection(conn)
	}
}
