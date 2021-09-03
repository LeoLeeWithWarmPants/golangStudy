package main

import "fmt"

func sum(n1 int, n2 int) int {
	//defer修饰的语句暂时不会执行，这些语句将会按照顺序被压入一个独立的栈中
	//当函数执行完毕后，按照栈先进后出的原则出栈并执行
	//在defer将语句压入栈中时，语句中相关变量的“值拷贝”也一同压入了栈中，所以defer的语句打印出来的n1和n2的值是1和2
	//defer的价值在于，当函数执行完毕后，进行一些资源关闭和回收的“善后”操作
	defer fmt.Printf("defer: n1=%d\n", n1)
	defer fmt.Printf("defer: n2=%d\n", n2)
	n1 += 2
	n2 += 2
	res := n1 + n2
	fmt.Printf("func sum res=%d\n", res)
	return res
}

func main() {
	res := sum(1, 2)
	fmt.Printf("func main res=%d\n", res)
}
