package main

import (
	"fmt"
	"net"
	"strconv"
)

func StartClient(address string, port int, retryTimes int) {
	var practicalRetryTimes int = 1
	var connection net.Conn
	var err error
	for {
		//重试
		connection, err = net.Dial("tcp", address+":"+strconv.FormatInt(int64(port), 10))
		if err != nil {
			fmt.Printf("connect times:%d, client connect failed, connect info:%v\n", practicalRetryTimes, connection.LocalAddr())
			if practicalRetryTimes >= retryTimes {
				fmt.Printf("retry %d times, give up\n", practicalRetryTimes)
				break
			}
			practicalRetryTimes++
		} else {
			fmt.Printf("client connect server success, remote:%v\n", connection.RemoteAddr())
			break
		}
	}
	if connection == nil || err != nil {
		//如果链接失败就结束程序
		return
	}

}

func main() {
	StartClient("127.0.0.1", 9528, 3)
}
