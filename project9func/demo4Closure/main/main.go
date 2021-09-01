package main

import "fmt"

// Add 闭包
func Add() func(int) int {
	var num int = 1
	//返回类型是一个函数类型，且是匿名函数，但这个匿名函数引用到了函数外的变量num，因此这个匿名函数就和num形成了一个整体，构成闭包
	//返回的函数类型私有化了变量num
	//所以在使用闭包时，要分析清楚返回函数引用到的变量有哪些，闭包=函数+该函数引用到的任何变量（包括外层函数的参数）
	return func(i int) int {
		num += i
		return num
	}
}

func main() {
	addFunc := Add()
	fmt.Printf("num=%d\n", addFunc(1)) //num=2
	fmt.Printf("num=%d\n", addFunc(2)) //num=4
	fmt.Printf("num=%d\n", addFunc(3)) //num=7
}
