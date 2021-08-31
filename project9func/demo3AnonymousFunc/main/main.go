package main

import (
	"fmt"
	"golangStudyProject/project9func/demo3AnonymousFunc/utils"
)

func main() {
	//定义匿名函数之后马上调用(一次性使用)
	res1 := func(num1 int8, num2 int8) int8 {
		fmt.Printf("calculate num1 + num2...\n")
		return num1 + num2
	}(1, 2)
	fmt.Printf("res1 = num1 + num2 = %d\n", res1)

	//将匿名函数赋值给变量，该变量的类型即为函数类型，可以重复使用
	res2 := func(num1 int8, num2 int8) int8 {
		fmt.Printf("calculate num1 + num2...\n")
		return num1 + num2
	}
	fmt.Printf("res2 = num1 + num2 = %d\n", res2(3, 4))

	//调用全局匿名函数
	res3 := utils.MyFunc(5, 6)
	fmt.Printf("res3 = num1 + num2 = %d\n", res3)
}
