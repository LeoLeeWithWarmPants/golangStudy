package main

import (
	"fmt"
	initUtil "golangStudyProject/project9func/demo2initFunc/utils"
)

//全局变量定义首先初始化，其次执行init函数，最后是main函数
var num = initNum()

func initNum() int {
	fmt.Println("func initNum executing...")
	return 20
}

//init函数总是在main函数之前执行，通常在init函数中进行一些初始化操作
func init() {
	fmt.Println("func init executing...")
	fmt.Printf("num=%d\n", num)
}

func main() {
	fmt.Println("func main executing...")
	fmt.Printf("num=%d\n", num)
	fmt.Printf("name=%q,age=%d\n", initUtil.Name, initUtil.Age)
}
