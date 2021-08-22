package main

import "fmt"

//指针
func main() {

	//基本数据类型的内存布局
	var i8 int8 = 10
	//i的内存地址
	fmt.Printf("i8的内存地址:%v,&i8的类型:%T\n", &i8, &i8)
	var ptr *int8 = &i8
	fmt.Printf("ptr的值:%v,ptr的内存地址:%v,ptr的类型:%T\n", ptr, &ptr, ptr)
	//根据指针类型获取真正的值,使用*来获取
	fmt.Printf("real value:%d\n", *ptr)
	//通过指针来访问、修改指向的内存空间
	*ptr = 11
	fmt.Printf("i8=%d\n", i8)
}
