package main

import "fmt"

type A interface {
	ATest()
}

type B interface {
	BTest()
}

//接口的继承，如果要实现C接口，那么就要将其继承的A、B接口的方法也一并实现
type C interface {
	A
	B
	CTest()
}

type Str1 struct {
}

func (s Str1) CTest() {
	fmt.Println("CTest...")
}

func (s Str1) BTest() {
	fmt.Println("BTest...")
}

func (s Str1) ATest() {
	fmt.Println("ATest...")
}

func InvokeCInterface(c C) {
	c.ATest()
	c.BTest()
	c.CTest()
}

func InvokeBInterface(b B) {
	b.BTest()
}

func main() {
	//接口继承测试
	s1 := new(Str1)
	InvokeCInterface(s1)
	InvokeBInterface(s1)
}
