package main

import (
	"fmt"
	"runtime"
	"time"
)

// 每隔疫苗输出一次指定文本
func printHelloWorld(times int, text string) {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Printf("%d.%s\n", i+1, text)
	}
}

//init函数始终在main函数执行之前执行
func init() {
	//获取当前环境的逻辑CPU核心数
	cpuNum := runtime.NumCPU()
	fmt.Printf("current logic-cpu num:%d\n", runtime.NumCPU())
	//可以自行设置使用多少个CPU
	var actualProcsNum int = 0
	if (cpuNum - 1) > 0 {
		actualProcsNum = cpuNum - 1
	}
	runtime.GOMAXPROCS(actualProcsNum)
	fmt.Printf("current logic-cpu num:%d\n", runtime.NumCPU())
}

func main() {
	//使用 go 关键字开启协程执行
	go printHelloWorld(10, "Hello world")

	printHelloWorld(10, "Hello LeoLee")
}
