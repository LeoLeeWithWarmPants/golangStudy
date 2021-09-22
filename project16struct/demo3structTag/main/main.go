package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Id   int64
	Name string
	Age  int8
}

//小写的字段名
type PersonDuplicate1 struct {
	id   int64
	name string
	age  int8
}

//tag方式定义字段序列化之后的名称
type PersonDuplicate2 struct {
	Id   int64  `json:"idddd"`
	Name string `json:"name"`
	Age  int8   `json:"age"`
}

func main() {
	p1 := new(Person)
	(*p1).Id = int64(1)
	(*p1).Name = "Leo Lee"
	(*p1).Age = 26

	//序列化struct对象为json格式字符串
	p1Json, error := json.Marshal(p1)
	if error != nil {
		fmt.Printf("p1 to json failed, error msg:%v\n", error)
	} else {
		//json.Marshal返回的是一个byte类型的切片（[]byte）
		fmt.Println("p1Json:", p1Json)
		//可以使用如下方式转为string字符串
		p1JsonStr := fmt.Sprintf("p1 to json:%s\n", p1Json)
		fmt.Println(p1JsonStr)
	}

	//如果struct中字段名为小写，代表私有字段，作用域为本包，所以在encoding/json包中无法使用，这时序列化将会忽略私有化的字段
	p2 := new(PersonDuplicate1)
	p2.id = int64(2)
	p2.name = "Lydia"
	p2.age = 29
	p2Json, error := json.Marshal(p2)
	if error != nil {
		fmt.Printf("p2 to json failed, error msg:%v\n", error)
	} else {
		fmt.Println("encoding/json中的函数无法使用私有的字段（作用域之外的字段），序列化自动忽略")
		fmt.Println("所以to Json操作字段名小写与作用域相冲突：字段名首字母小写就无法被包外函数所使用")
		fmt.Println("p2Json:", p2Json)
		p2JsonStr := fmt.Sprintf("p2 to json:%s\n", p2Json)
		fmt.Println(p2JsonStr)
	}

	//怎么让序列化之后的json字符串中的字段名首字母为小写呢？？？
	//使用tag
	p3 := &PersonDuplicate2{
		Id:   3,
		Name: "Tony",
		Age:  45,
	}
	p3Json, error := json.Marshal(p3)
	if error != nil {
		fmt.Printf("p3 to json failed, error msg:%v\n", error)
	} else {
		fmt.Println("使用tag的方式修饰字段，序列化之后的字段名首字母为小写")
		p3JsonStr := fmt.Sprintf("p3 to json:%s\n", p3Json)
		fmt.Println(p3JsonStr)
	}
}
