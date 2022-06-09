package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
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
	//发送数据到服务端
	terminalStr := terminalInputRead()
	//发送数据到服务端
	_, writeErr := connection.Write([]byte(terminalStr))
	if writeErr != nil {
		fmt.Printf("write to server failed, root of problem is Conn.Write failed, err=%v\n", writeErr)
	}
}

func terminalInputRead() string {
	terminalReader := bufio.NewReader(os.Stdin) //标准输入
	inputStr, readErr := terminalReader.ReadString('\n')
	if readErr != nil {
		fmt.Printf("terminal input read failed, err=%v\n", readErr)
	}
	return inputStr
}

func main() {
	StartClient("127.0.0.1", 9528, 3)
}
