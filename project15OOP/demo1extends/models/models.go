package models

import "fmt"

type Person struct {
	Name string
	Age  int
	Book Book //有名结构体，访问其字段的时候需要使用其名称作出具体的调用，如：Person.book.Name
	Phone
}

func (p *Person) ShowPersonInfo() {
	fmt.Printf("Name:%s,Age:%d\n", (*p).Name, (*p).Age)
}

type Student struct {
	Person //使用嵌套匿名结构体的方式来实现继承
	score  int
}

func (stu *Student) getScore() int {
	return stu.score
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

type Book struct {
	Name  string
	Price float64
}

type Phone struct {
	Price float64
}

//在同一个package中，struct可以使用嵌套struct中所有的字段和方法，不论是首字母大写还是小写。
func GetScore1(undergraduate *Undergraduate) int {
	return undergraduate.score
}

func GetScore2(undergraduate *Undergraduate) int {
	return undergraduate.Student.getScore()
}
