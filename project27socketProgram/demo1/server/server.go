package main

import (
	"fmt"
	"io"
	"net"
	"strconv"
)

func StartServer(protocol, address string, port int) {
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
		//读取数据
		go msgProcessor(conn)
	}
}

func msgProcessor(conn net.Conn) {
	//链接使用完一定要关闭
	defer conn.Close()

	for {
		//每次读取创建一个新的缓冲区
		readBuf := make([]byte, 1024)
		readNum, readErr := conn.Read(readBuf) // 如果客户端没有发消息，read将会在超时时间到来之前一只阻塞
		if readErr != nil {
			if readErr == io.EOF {
				fmt.Printf("remote client was closed, client address=%v\n", conn.RemoteAddr())
			} else {
				fmt.Printf("server read faild, client address=%v, err=%v\n", conn.RemoteAddr(), readErr)
			}
			break
		}
		//打印信息
		fmt.Print(string(readBuf[:readNum]))
	}
}

func main() {
	StartServer("tcp", "0.0.0.0", 9528)
	//启动后telnet 127.0.0.1 9528
}
