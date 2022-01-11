package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

type Student struct {
	BaseInfo Person `json:"baseInfo"`
	Type     int    // 1-小学生 2-中学生 3-大学生
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

func fromJson(jsonStr string, v interface{}) {
	err := json.Unmarshal([]byte(jsonStr), &v)
	if err != nil {
		fmt.Printf("deserialization failed, err:%v\n", err)
	}
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
	fmt.Printf("slice1=%v\n", slice1)
	toJson(slice1)

	//反序列化
	slice2 := make([]interface{}, 2)
	//由于slice2在声明的时候，类型为空接口，json.Unmarshal并不能确定其中元素的实际类型，如果不在这里初始化slice的类型，反序列化之后，slice2的元素将会都是map类型
	//这里属于显式的声明slice中元素的类型，使序列化的时候按照当前类型进行序列化
	slice2[0] = make(map[string]interface{})
	slice2[1] = new(Student)
	fromJson("[{\"info\":{\"name\":\"LeoLee\",\"age\":18},\"name\":\"LeoLee\"},{\"baseInfo\":{\"name\":\"LeoLee\",\"age\":18},\"Type\":3,\"Grade\":\"大一\"}]", &slice2)
	fmt.Printf("deserialization success, slice2=%v\n", slice2)
}
