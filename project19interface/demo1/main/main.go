package main

import (
	"fmt"
	"golangStudyProject/project15OOP/demo1extends/models"
)

//定义接口
type PersonInterface interface {
	//定义接口内方法
	TakeMetro() //坐地铁
	DoWork()    //工作
	Study()     //学习
}

type DemoInterface interface {
	AutoIncrement(i int) int
}

type Demo struct {
}

func (d Demo) AutoIncrement(i int) int {
	i++
	return i
}

//定义结构体Student
type Student struct {
	models.Person
}

type Programmer struct {
	models.Person
}

//Student实现PersonInterface定义的方法
func (stu Student) TakeMetro() {
	fmt.Printf("学生[%s]坐地铁了\n", stu.Name)
}

func (stu Student) DoWork() {
	fmt.Printf("学生[%s]正在做兼职\n", stu.Name)
}

func (stu Student) Study() {
	fmt.Printf("学生[%s]正在学习\n", stu.Name)
}

//Programmer实现PersonInterface定义的方法
func (pro Programmer) TakeMetro() {
	fmt.Printf("程序员[%s]坐地铁了\n", pro.Name)
}

func (pro Programmer) DoWork() {
	fmt.Printf("程序员[%s]正在搬砖\n", pro.Name)
}

func (pro Programmer) Study() {
	fmt.Printf("程序员[%s]正在学习\n", pro.Name)
}

//调用接口的方法
func PersonDoSomething(personInterface PersonInterface) {
	personInterface.TakeMetro()
	personInterface.DoWork()
	personInterface.Study()
}

func DemoInterfaceTest(demoInterface DemoInterface, i int) int {
	resultInt := demoInterface.AutoIncrement(i)
	return resultInt
}

func main() {
	p1 := models.Person{Name: "Lydia"}
	p2 := models.Person{Name: "LeoLee"}
	stu := &Student{p1}
	pro := &Programmer{p2}
	//将接口的实现传入函数中，函数中调用接口的具体实现方法
	PersonDoSomething(stu)
	PersonDoSomething(pro)

	//接口定义带参有返回值
	d := new(Demo)
	i := DemoInterfaceTest(d, 2)
	fmt.Printf("i=%d\n", i)
}
