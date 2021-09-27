package main

import (
	"fmt"
	"golangStudyProject/project15OOP/demo1extends/models"
)

//父struct的方法，可以被子struct使用
func main() {
	p1 := &models.Person{Name: "WangWu", Age: 18}
	p1.ShowPersonInfo()

	stu1 := &models.Student{}
	stu1.Person.Name = "ZhangSan"
	stu1.Person.Age = 6
	stu1.SetStuScore(80)
	stu1.ShowPersonInfo()
	stu1.ShowStuInfo()

	mss1 := new(models.MiddleSchoolStudent)
	(*mss1).Student.Name = "LeoLee"
	(*mss1).Student.Age = 26
	mss1.Test()
	mss1.SetStuScore(99)
	stu1.ShowPersonInfo()
	mss1.ShowStuInfo()

	u1 := new(models.Undergraduate)
	(*u1).Student.Name = "Lydia"
	(*u1).Student.Age = 29
	u1.Test()
	u1.SetStuScore(90)
	stu1.ShowPersonInfo()
	u1.ShowStuInfo()

	//在同一个package中，struct可以使用嵌套struct中所有的字段和方法，不论是首字母大写还是小写。
	mss2 := new(models.Undergraduate)
	mss2.SetStuScore(70)
	var score int
	score = models.GetScore1(mss2)
	fmt.Println("score=", score)
	mss2.SetStuScore(60)
	score = models.GetScore2(mss2)
	fmt.Println("score=", score)
}
