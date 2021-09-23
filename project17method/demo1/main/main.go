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

//使用指针进行引用传递
func (a *A) Amethod3(i int) {
	fmt.Printf("a type:%T,a address:%p\n", a, a)
	(*a).Id += i //golang由于编译器优化，也可以写为：a.Id += i
}

//自定义String()方法
func (a *A) String() string {
	return fmt.Sprintf("自定义String()方法执行，a=%v\n", *a)
}

func main() {
	a1 := &A{Id: 8}
	a1.Amethod() //struct A的变量调用其自有的方法Amethod，会把变量自身作为参数传递给Amethod

	a1.Amethod2(9)
	fmt.Printf("a1:%v\n", *a1) //可以看出：对象调用方法的时候，会将自己的副本作为参数进行传递

	(*a1).Amethod3(2) //使用指针作为方法的参数，就可以进行引用传递，golang由于编译器优化，也可以写为：a1.Amethod3(2)
	fmt.Printf("a1:%v, a1 address:%p\n", *a1, a1)

	//如果一个类型实现了一个String()这个方法，那么fmt.Println默认会调用这个类型的这个String()方法，相当于Java重写了toString()
	a1str := (*a1).String()
	fmt.Printf("a1str=%s\n", a1str)
	fmt.Printf("a1 type:%T\n", a1) //a1为指针类型
	fmt.Println(a1)                //如果传递是指针，则调用了自定义的String()方法
	fmt.Println(*a1)
	a2 := A{Id: 2}
	fmt.Printf("a2 type:%T\n", a2)
	fmt.Println(a2)
	fmt.Println(&a2) //如果传递是指针，则调用了自定义的String()方法
}
