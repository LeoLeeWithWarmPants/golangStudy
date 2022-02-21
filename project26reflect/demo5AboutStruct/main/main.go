package main

import (
	"fmt"
	"reflect"
)

type A struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

//使用指针进行引用传递
func (a *A) Amethod2(i int) {
	fmt.Printf("a type:%T,a address:%p\n", a, a)
	(*a).Id += i //golang由于编译器优化，也可以写为：a.Id += i
}

func (a A) Amethod3(i int) bool {
	if a.Id > i {
		return true
	}
	return false
}

//定义一个属于struct A的方法
func (a A) Amethod() A {
	fmt.Println("这是一个属于struct A的方法，名为：Amethod")
	fmt.Printf("param a:%v\n", a)
	return a
}

//自定义String()方法
func (a *A) String() string {
	return fmt.Sprintf("自定义String()方法执行，a=%v\n", *a)
}

func reflectStruct(s interface{}) {
	rValue := reflect.ValueOf(s)
	rType := reflect.TypeOf(s)
	if rValue.Kind() != reflect.Struct {
		fmt.Println("can't support this type, only Struct!")
		return
	}
	//获取结构体所有字段
	fieldNum := rValue.NumField()
	fmt.Printf("this Struct include %d fieldes\n", fieldNum)
	for i := 0; i < fieldNum; i++ {
		currentField := rValue.Field(i)
		currentStructField := rType.Field(i)
		fmt.Printf("%d field[%s]的值为: %v, 类型为: %v, tag: %s\n", i, currentStructField.Name, currentField, currentField.Kind(), currentStructField.Tag.Get("json"))
	}
	fmt.Println("")
	//获取结构体方法
	methodNum := rValue.NumMethod()
	fmt.Printf("this Struct include %d methodes\n", methodNum)
	fmt.Println("rValue.NumMethod()只统计结构体对象的方法，结构体指针的方法不参与统计")
	//调用第二个方法
	rValue.Method(0).Call(nil) //调用index为0的方法，也就是Amethod
	//为什么Amethod3在Amethod方法前面，0代表的却是Amethod呢？？？
	//因为方法的排序默认是按照函数名的ASCII码排序的
	//调用带参方法
	var params []reflect.Value
	params = append(params, reflect.ValueOf(11))
	b := rValue.Method(1).Call(params)
	//类型断言
	switch b[0].Kind() {
	case reflect.Bool:
		fmt.Printf("Amethod3 result:%t\n", b[0].Bool())
	default:
		fmt.Println("暂不支持解析的数据类型")
	}
}

func main() {
	a := A{Id: 10, Name: "LeoLee"}
	reflectStruct(a) //如果传递a的指针，还可以通过反射修改字段值
}
