package main

import (
	"fmt"
	"time"
)

//协程中的panic将会导致程序崩溃，需要进行人工处理，防止某一个协程的错误影响到整个程序

func rightFunc() {
	for i := 0; i < 10; i++ {
		time.Sleep(time.Second)
		fmt.Println("正常执行：", i)
	}
}

func errorFunc() {
	defer catchPanic()
	var map1 map[string]interface{}
	//这里没有make，将会报错，panic，终止程序运行
	map1["a"] = "LeoLee"
}

//捕获异常函数
func catchPanic() {
	err := recover()
	if err != nil {
		fmt.Printf("发生异常：err=%v\n", err)
	}
}

//errorFunc中不做任何处理，errorFunc的panic将会终止程序运行
func process() {
	go rightFunc()
	time.Sleep(5 * time.Second)
	go errorFunc()
	//为了让主线程不马上执行完成
	time.Sleep(12 * time.Second)
}
func main() {
	fmt.Println("test start")
	process()
	fmt.Println("test end")
}
