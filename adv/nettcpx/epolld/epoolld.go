package main

import (
	"fmt"
	"net"
)

// func main() {
// 	// 1. 创建 epoll 实例
// 	epfd, err := syscall.EpollCreate1(0)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	// 2. 创建监听套接字
// 	listener, err := net.Listen("tcp", "127.0.0.1:8080")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer listener.Close()
//
// 	fd := int(listener.(*net.TCPListener).File().Fd())
//
// 	// 3. 将监听套接字添加到 epoll 实例中
// 	event := syscall.EpollEvent{
// 		Events: syscall.EPOLLIN,
// 		Fd:     int32(fd),
// 	}
// 	err = syscall.EpollCtl(epfd, syscall.EPOLL_CTL_ADD, fd, &event)
// 	if err != nil {
// 		log.Fatal(err)
// 	}
//
// 	events := make([]syscall.EpollEvent, 10)
//
// 	for {
// 		// 4. 循环等待EpollWait返回，有返回表示有FD就绪
// 		n, err := syscall.EpollWait(epfd, events, -1)
// 		if err != nil {
// 			log.Fatal(err)
// 		}
//
// 		// 5. 开始从Epoll就绪列表获取就绪FD进行IO事件读\写处理
// 		for i := 0; i < n; i++ {
// 			if int(events[i].Fd) == fd { // 当前就绪到为监听套接字
//
// 				// 6. 具体连接就绪IO事件读取返回，获取到连接套接字
// 				conn, err := listener.Accept()
// 				if err != nil {
// 					log.Println("Error accepting connection:", err)
// 					continue
// 				}
//
// 				// 7. 针对连接套接字进行业务操作
// 				go handleConnection(conn)
// 			}
// 		}
// 	}
// }

func handleConnection(conn net.Conn) {
	defer conn.Close()
	fmt.Println("Accepted connection from", conn.RemoteAddr())

	// 在这里处理连接的读写逻辑
}
