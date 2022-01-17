package main

import (
	"fmt"
	"time"
)

// 每隔疫苗输出一次指定文本
func printHelloWorld(times int, text string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%d.%s\n", i+1, text)
	}
}

func main() {
	//使用 go 关键字开启协程执行
	go printHelloWorld(10, "Hello world")

	printHelloWorld(10, "Hello LeoLee")
}
