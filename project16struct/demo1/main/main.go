package main

import "fmt"

//struct定义
type Person struct {
	Id   int64
	Name string
	Age  int8
}

type Student struct {
	baseInfo Person
	ptr      *int //指针
	array    [2]int
	slice    []int
	map1     map[string]int
}

func changePerson(p Person) {
	p.Name = "Lydia"
}

func main() {
	//定义Person类型的变量stu1
	var p1 Person
	fmt.Printf("p1.Id=%d,p1.Name=%s,p1.Age=%d\n", p1.Id, p1.Name, p1.Age) //初始化变量完成后，各个字段为类型的默认值
	p1.Id = int64(1000)
	p1.Name = "LeoLee"
	p1.Age = 26
	fmt.Printf("p1.Id=%d,p1.Name=%s,p1.Age=%d\n", p1.Id, p1.Name, p1.Age)
	//struct是值类型，函数内的修改无法影响函数外的值。
	//需要进行指针才可以改变原有变量的数据
	changePerson(p1)
	fmt.Printf("p1.Id=%d,p1.Name=%s,p1.Age=%d\n", p1.Id, p1.Name, p1.Age)
	p2 := p1
	p2.Name = "Lydia"
	fmt.Printf("p1.Id=%d,p1.Name=%s,p1.Age=%d\n", p1.Id, p1.Name, p1.Age)
	fmt.Printf("p2.Id=%d,p2.Name=%s,p2.Age=%d\n", p2.Id, p2.Name, p2.Age)
	p3 := &p1
	p3.Name = "Tony"
	fmt.Printf("p1.Id=%d,p1.Name=%s,p1.Age=%d\n", p1.Id, p1.Name, p1.Age)
	fmt.Printf("p3.Id=%d,p3.Name=%s,p3.Age=%d\n", p3.Id, p3.Name, p3.Age)

	var stu1 Student
	fmt.Println(stu1)
	//prt、slice、map都为nil，prt是因为没有指向内存，slice和map是因为没有make
	if stu1.ptr == nil {
		fmt.Printf("stu1.ptr=%v\n", stu1.ptr)
		//指针类型赋值
		i := 8
		stu1.ptr = &i
		fmt.Printf("stu1.ptr[%p], value=%d\n", stu1.ptr, *stu1.ptr)
	}
	if stu1.slice == nil {
		fmt.Printf("stu1.slice=%v\n", stu1.slice)
		//初始化slice
		stu1.slice = make([]int, 2)
		stu1.slice[0] = 1
		stu1.slice[1] = 2
		stu1.slice = append(stu1.slice, 3)
	}
	if stu1.map1 == nil {
		fmt.Printf("stu1.map1=%v\n", stu1.map1)
		//初始化map
		stu1.map1 = make(map[string]int)
		stu1.map1["A"] = 5
		stu1.map1["B"] = 6
	}
	stu1.baseInfo.Name = "LeoLee"
	fmt.Println(stu1)
}