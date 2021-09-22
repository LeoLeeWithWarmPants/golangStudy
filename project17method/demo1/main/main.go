package main

import "fmt"

type A struct {
	Id int
}

//定义一个属于struct A的方法
func (a A) Amethod() A {
	fmt.Println("这是一个属于struct A的方法，名为：Amethod")
	fmt.Printf("param a:%v\n", a)
	return a
}

func (a A) Amethod2(i int) A {
	a.Id = i
	fmt.Printf("Amethod2.a:%v\n", a)
	return a
}

func main() {
	a1 := &A{Id: 8}
	a1.Amethod() //struct A的变量调用其自有的方法Amethod，会把变量自身作为参数传递给Amethod

	a1.Amethod2(9)
	fmt.Printf("a1:%v\n", a1) //可以看出：对象调用方法的时候，会将自己的副本作为参数进行传递
}
