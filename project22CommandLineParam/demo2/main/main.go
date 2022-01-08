package main

import (
	"flag"
	"fmt"
)

var userName string
var pwd string
var host string
var port int

func main() {
	flag.StringVar(&userName, "u", "LeoLee", "用户名")
	flag.StringVar(&pwd, "p", "123456", "密码")
	flag.StringVar(&host, "h", "127.0.0.1", "地址")
	flag.IntVar(&port, "port", 3306, "端口")

	//重要操作
	flag.Parse()

	fmt.Printf("%s %s %s %d\n", userName, pwd, host, port)
}
