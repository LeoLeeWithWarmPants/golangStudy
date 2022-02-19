package main

import "fmt"

func main() {
	const age int = 27
	const name = "LeoLee" // 常量的类型可写可不写

	//简洁声明
	const (
		a int = 1
		b     = 2
		c     = true
	)

	//iota, 表示给d赋值0, 之后的e f g将会在d的基础上+1
	//iota本身就是常量 const iota = 0
	const (
		d = iota
		e
		f
		g, h = iota, iota // 此时g == h == 3，同一行的iota不会+1
	)
	fmt.Println(d, e, f, g, h)
}
