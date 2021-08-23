package main

import "fmt"

//本阶段为计算机语言基础学习，详见readme.md
func main() {
	//算数运算符中需要特别注意除法，与Java一样，此处结果为2，省略了小数部分
	fmt.Println("10/4 = ", 10/4)
	//Java中float f = 10/4;f的结果是2.0，与Java一样，都不会是2.5，Golang中结果是2.000000
	var f float32 = 10 / 4
	fmt.Printf("10/4 = %f\n", f)
	//想要输出预期的小数，需要保证有浮点类型参与运算
	fmt.Printf("10/4 = %g\n", 10.0/4)
	fmt.Printf("10/4 = %g\n", 10/4.0)
	fmt.Printf("10/4 = %g\n", float32(10)/4)

	//++ --在Golang中只能独立使用，否则编译报错
	var i8 int8 = 2
	//不能直接使用 a := i8++
	i8++
	a := i8
	fmt.Printf("a=%d\n", a)
	//在Golang中没有 ++i --i
	//甚至在if中都不可以使用
	/*if i8++ > 0 {
		// do something
	}*/

	//多重赋值
	var c = 20
	var b = 30
	c, b = b, c
	fmt.Printf("c=%d, b=%d\n", c, b)

	//二进制
	var i int8 = 10
	fmt.Printf("i的二进制:%b\n", i)
	//八进制
	var x int8 = 012
	fmt.Printf("x=%d\n", x)
	//十六进制
	var y int8 = 0x11
	fmt.Printf("y=%d\n", y)
}
