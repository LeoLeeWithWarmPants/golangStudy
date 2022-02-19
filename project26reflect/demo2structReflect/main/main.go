package main

import (
	"fmt"
	"reflect"
)

type person struct {
	Name string
	Age  int
}

func reflectStruct(s interface{}) {
	//获取reflect.Type
	reflectType1 := reflect.TypeOf(s)
	fmt.Printf("reflectType1:%v\n", reflectType1)
	fmt.Printf("reflectType1 type:%T\n", reflectType1)

	//获取reflect.Value
	reflectValue1 := reflect.ValueOf(s)
	fmt.Printf("reflectValue1:%v\n", reflectValue1)
	fmt.Printf("reflectValue1 type:%T\n", reflectValue1)

	//reflect.Value to interface{}
	interface1 := reflectValue1.Interface()
	p, ok := interface1.(person)
	fmt.Printf("type assertion:%t, interface1 value:%v\n", ok, p)
}

func main() {
	p1 := person{Name: "LeoLee", Age: 27}
	reflectStruct(p1)
}
