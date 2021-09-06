package main

import "fmt"

func test() {
	//使用defer + recover来捕获和处理异常
	defer func() {
		err := recover() //内建函数，捕获异常
		if err != nil {
			fmt.Printf("err=%v/n", err)
		}
	}()
	//产生运行时异常
	n1 := 1
	n2 := 0
	fmt.Printf("result=%d\n", n1/n2)
}

//异常处理
func main() {
	test()
	fmt.Println("main函数执行结束")
}
