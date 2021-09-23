package main

import "fmt"

type Person struct {
	Id   int64
	Name string
	Age  int8
}

//打印Person
func (p Person) PrintPerson() {
	fmt.Printf("Id:%d, Name:%s, Age:%d\n", p.Id, p.Name, p.Age)
}

//增加年纪
func (p Person) AddAge(age int8) Person {
	p.Age += age
	return p
}

func main() {
	p1 := Person{Id: 20001, Name: "LeoLee", Age: 26}
	p1.PrintPerson()

	p1 = p1.AddAge(2)
	p1.PrintPerson()
}
