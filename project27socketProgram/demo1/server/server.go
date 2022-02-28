package main

import (
	"fmt"
	"net"
	"strconv"
)

func startServer(protocol, address string, port int) {
	listener, error := net.Listen(protocol, address+":"+strconv.FormatInt(int64(port), 10))
	defer listener.Close()
	if error != nil {
		fmt.Printf("listen error:%v\n", error)
		panic(error)
	}
	fmt.Printf("listen successful, wait for client connect\n")
	for {
		//这里等待客户端链接会阻塞的，直到有客户端进行链接
		conn, error := listener.Accept()
		if error != nil {
			fmt.Printf("accept error:%v\n", error)
			fmt.Printf("remote address:%v\n", conn.RemoteAddr())
		} else {
			fmt.Printf("remote address[%v] connected success\n", conn.RemoteAddr())
		}
	}
}

func main() {
	startServer("tcp", "0.0.0.0", 9528)
	//启动后telnet 127.0.0.1 9528
}
