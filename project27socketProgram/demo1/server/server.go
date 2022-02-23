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
	for {
		conn, error := listener.Accept()
		if error != nil {
			fmt.Printf("accept error:%v\n", error)
			fmt.Printf("remote address:%v\n", conn.RemoteAddr())
		} else {
			fmt.Printf("remote address[%v] connected success\n", conn.RemoteAddr())
		}
	}
}
