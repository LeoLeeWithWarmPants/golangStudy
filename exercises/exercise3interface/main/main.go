package main

import (
	"fmt"
	"math/rand"
	"sort"
)

//定义结构体
type Person struct {
	Name string
}

type Student struct {
	PersonInfo Person
	Score      int
}

type StudentSlice []Student

//实现sort中的Interface接口
func (studentSlice StudentSlice) Len() int {
	return len(studentSlice)
}
func (studentSlice StudentSlice) Less(i, j int) bool {
	//定义排序策略：降序
	return studentSlice[i].Score > studentSlice[j].Score
}
func (studentSlice StudentSlice) Swap(i, j int) {
	temp := studentSlice[i]
	studentSlice[i] = studentSlice[j]
	studentSlice[j] = temp
}

//对结构体进行排序
func main() {
	var students StudentSlice
	for i := 0; i < 10; i++ {
		student := Student{
			PersonInfo: Person{
				Name: fmt.Sprintf("LeoLee%d", i),
			},
			Score: rand.Intn(100),
		}
		students = append(students, student)
	}

	fmt.Println("排序前：")
	for _, e := range students {
		fmt.Println(e)
	}

	sort.Sort(students)
	fmt.Println("排序后：")
	for _, e := range students {
		fmt.Println(e)
	}
}
