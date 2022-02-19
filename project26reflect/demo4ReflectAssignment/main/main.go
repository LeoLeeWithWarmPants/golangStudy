package main

import (
	"fmt"
	"reflect"
)

func reflectSomething(s interface{}) {
	rValue := reflect.ValueOf(s)
	rKind := rValue.Kind()
	fmt.Printf("rKind = %v, rKind type = %T\n", rKind, rKind)
	rType := reflect.TypeOf(s)
	fmt.Printf("rType:%v\n", rType)
	fmt.Printf("rType type:%T\n", rType)
	fmt.Printf("rType Name:%v\n", rType.Name())
	switch rType.String() {
	case "*int":
		//通过Value的Elem方法获取指针的值
		rValue.Elem().SetInt(rValue.Elem().Int() + 1)
	default:
		fmt.Println("暂不支持解析的数据类型")
	}
}

//反射之后的赋值操作
func main() {
	var num1 int = 0
	reflectSomething(&num1)
	fmt.Printf("num1 = %v\n", num1)
}
