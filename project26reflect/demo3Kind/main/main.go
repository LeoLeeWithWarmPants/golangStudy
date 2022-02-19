package main

import "fmt"
import "reflect"

//Kind可以理解为某种类型的常量表示方式
func reflectSomething(s interface{}) {
	reflectType1 := reflect.TypeOf(s)
	fmt.Printf("reflectType1:%v\n", reflectType1)
	fmt.Printf("reflectType1 type:%T\n", reflectType1)
	reflectValue1 := reflect.ValueOf(s)
	fmt.Printf("reflectValue1:%v\n", reflectValue1)
	fmt.Printf("reflectValue1 type:%T\n", reflectValue1)

	//获取Kind
	kind1 := reflectType1.Kind()
	fmt.Printf("kind1 = %v, kind1 type = %T\n", kind1, kind1)
	kind2 := reflectValue1.Kind()
	fmt.Printf("kind2 = %v, kind2 type = %T\n", kind2, kind2)
}

func main() {
	var num1 int
	num1 = 2
	reflectSomething(num1)
}
