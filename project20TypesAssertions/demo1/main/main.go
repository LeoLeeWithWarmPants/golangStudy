package main

import "fmt"

type Coordinate struct {
	x int
	y int
}

//空接口，任何类型都可以直接实现了空接口
type emptyInterface interface {
}

//引出类型断言的概念demo
func main() {
	c1 := Coordinate{1, 2}
	var e emptyInterface
	e = c1 //OK

	c2 := Coordinate{2, 3}
	// c2 = e 编译错误：Cannot use 'e' (type emptyInterface) as the type Coordinate
	c2 = e.(Coordinate) //类型断言，类似于Java中的类型强制转换，如果变量e不是Coordinate类型，那将会产生恐慌
	fmt.Println(c2)

	//可检测的类型断言
	var e2 emptyInterface
	var f1 float64 = 1.1
	e2 = f1
	y, isOk := e2.(float64)
	fmt.Printf("e2类型断言：%t，", isOk)
	fmt.Printf("y的类型是%T\n", y)
}
