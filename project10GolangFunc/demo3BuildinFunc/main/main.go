package main

import "fmt"

//TODO 内建函数
//golang的设计者，为了方便使用，提供了一些内建函数，这些函数可以直接使用，所有内建函数api：https://studygolang.com/pkgdoc
func main() {
	//new
	n1 := 1
	fmt.Printf("n1=%v, n1的类型=%T, n1的地址%v\n", n1, n1, &n1)
	n2 := new(int) // n2的类型是指针，n2的值是一个内存地址，同时n2也有自己的地址，这个地址在这里指向的是int类型默认值0的地址
	fmt.Printf("n2=%v, n2的类型=%T, n2的地址%v\n", n2, n2, &n2)
	fmt.Printf("n2这个指针类型的变量指向的实际地址的值:%v\n", *n2)
	//修改n2的实际值
	*n2 = 20
	fmt.Printf("n2这个指针类型的变量指向的实际地址的值:%v\n", *n2)

}
