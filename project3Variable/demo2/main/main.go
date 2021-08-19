package main

import "fmt"

//全局变量
var n1 int = 1
var n2 = 2

//全局变量不可以使用n3 := 3方式省略var关键字

//批量声明
var (
	n3   = 100
	n4   = 200
	str1 = "Leo Lee"
)

//变量的使用
func main() {
	//1.指定变量类型，声明变量后不赋值，变量将会使用默认值
	var i int //默认值为0
	fmt.Printf("i=%d\n", i)

	//2.类型推导：根据值自行推断变量的类型
	var decimals = 3.1415926
	fmt.Printf("decimals=%e\n", decimals)
	fmt.Printf("decimals=%f\n", decimals)
	fmt.Printf("decimals=%E\n", decimals)
	fmt.Printf("decimals=%F\n", decimals)
	fmt.Printf("decimals=%g\n", decimals)
	fmt.Printf("decimals=%G\n", decimals)
	var str = "This is a character string"
	fmt.Printf("%s\n", str)

	//3.使用 := 省略var，左侧的变量名不能是已经声明过的，否则编译失败
	name := "Leo Lee" //等价于 var name string = "Leo Lee"
	fmt.Printf("My name is %s\n", name)

	//4.多变量声明
	var v1, v2, v3 int //都是默认值0
	fmt.Printf("v1 = %d, v2 = %d, v3 = %d\n", v1, v2, v3)
	//声明多变量并赋值，同样可以类型推导
	var num, name2, decimals2 = 10, "Leo Lee", 26.5
	fmt.Printf("num = %d, name = %s, decimals = %f\n", num, name2, decimals2)
	num1, name3, decimals3 := 10, "Leo Lee", 26.5
	fmt.Printf("num = %d, name = %s, decimals = %f\n", num1, name3, decimals3)

	//5.全局变量
	fmt.Printf("n1 = %d, n2 = %d\n", n1, n2)
	fmt.Printf("n3 = %d, n4 = %d, str1 = %s\n", n3, n4, str1)
}
