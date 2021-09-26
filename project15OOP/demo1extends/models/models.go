package models

import "fmt"

type Person struct {
	Name string
	Age  int
}

func (p *Person) ShowPersonInfo() {
	fmt.Printf("Name:%s,Age:%d\n", (*p).Name, (*p).Age)
}

type Student struct {
	Person //使用嵌套匿名结构体的方式来实现继承
	score  int
}

// ShowStuInfo 打印学生信息
func (stu *Student) ShowStuInfo() {
	fmt.Printf("stu.Name:%s,stu.Age:%d,stu.Score:%d\n", (*stu).Name, (*stu).Age, (*stu).score)
}

// SetStuScore 设置学生成绩
func (stu *Student) SetStuScore(score int) {
	(*stu).score = score
}

type MiddleSchoolStudent struct {
	Student //使用嵌套匿名结构体的方式来实现继承
}

func (stu *MiddleSchoolStudent) Test() {
	fmt.Println("中学生考试...")
}

type Undergraduate struct {
	Student //使用嵌套匿名结构体的方式来实现继承
}

func (stu *Undergraduate) Test() {
	fmt.Println("大学生考试...")
}
