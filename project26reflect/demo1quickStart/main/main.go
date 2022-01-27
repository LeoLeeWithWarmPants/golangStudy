package main

import (
	"fmt"
	"reflect"
)

func baseReflect(s interface{}) {
	//获取reflect.Type
	reflectType1 := reflect.TypeOf(s)
	fmt.Printf("reflectType1:%v\n", reflectType1)
	fmt.Printf("reflectType1 type:%T\n", reflectType1)

	//获取reflect.Value
	reflectValue1 := reflect.ValueOf(s)
	fmt.Printf("reflectValue1:%v\n", reflectValue1)
	fmt.Printf("reflectValue1 type:%T\n", reflectValue1)

	//从reflect.Value中获取实际的值
	realValue1 := reflectValue1.Int()
	fmt.Printf("reflectValue1 real value：%d, type:%T\n", realValue1, reflectValue1.Int())

	//reflect.Value to interface{}
	interface1 := reflectValue1.Interface()
	num1, ok := interface1.(int)
	fmt.Printf("type assertion:%t, interface1 value:%d\n", ok, num1)
}

func main() {
	//对基本数据类型、interface{}、reflect.Value进行反射操作
	var num1 int = 1
	baseReflect(num1)
}
