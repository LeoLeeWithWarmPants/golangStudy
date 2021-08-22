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
}
