package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

type Student struct {
	BaseInfo Person
	Type     int // 1-小学生 2-中学生 3-大学生
	Grade    string
}

func toJson(v interface{}) string {
	jsonByte, err := json.Marshal(v)
	if err != nil {
		fmt.Printf("serialization failed, err:%v\n", err)
	}
	fmt.Printf("json result:%s\n", string(jsonByte))
	return string(jsonByte)
}

func main() {
	//结构体
	person := Person{Name: "LeoLee", Age: 18}
	student := new(Student)
	student.BaseInfo = person
	student.Type = 3
	student.Grade = "大一"
	toJson(student)

	//map
	map1 := make(map[string]interface{})
	map1["name"] = "LeoLee"
	map1["info"] = student
	toJson(map1)

	//切片
	map2 := make(map[string]interface{})
	map2["name"] = "LeoLee"
	map2["info"] = person

	var slice1 []interface{} = make([]interface{}, 2, 2)
	slice1[0] = map2
	slice1[1] = student
	toJson(slice1)
}
