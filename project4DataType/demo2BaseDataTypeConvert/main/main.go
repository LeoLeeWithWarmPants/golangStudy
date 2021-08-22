package main

import (
	"fmt"
	"strconv"
)

//基本数据类型的转换
func main() {
	var i8 int8 = 100
	var f32 float32 = float32(i8)
	fmt.Printf("f32=%f\n", f32)
	var i32 int32 = int32(i8)
	fmt.Printf("i32=%d\n", i32)
	//int64转int8，编译不报错，转换的结果会按照溢出处理
	var i64 int64 = 123123123
	i8 = int8(i64)
	fmt.Printf("i8=%d\n", i8)

	//基本类型转string
	var num1 int8 = 26
	var num2 uint8 = 95
	var num3 float64 = 3.1415926
	b := true
	var char byte = 'L'
	var str string
	//fmt.Sprintf方法
	str = fmt.Sprintf("%d", num1)
	fmt.Printf("str=%s\n", str)
	str = fmt.Sprintf("%d", num2)
	fmt.Printf("str=%s\n", str)
	str = fmt.Sprintf("%G", num3)
	fmt.Printf("str=%s\n", str)
	str = fmt.Sprintf("%t", b)
	fmt.Printf("str=%s\n", str)
	str = fmt.Sprintf("%c", char)
	fmt.Printf("str=%s\n", str)
	//strconv函数
	str = strconv.FormatInt(int64(num1), 10)
	fmt.Printf("str=%s\n", str)
	str = strconv.Itoa(int(num2))
	fmt.Printf("str=%s\n", str)
	str = strconv.FormatFloat(num3, 'f', -1, 64)
	fmt.Printf("str=%s\n", str)
	str = strconv.FormatBool(b)
	fmt.Printf("str=%s\n", str)

	//string转基本类型
	str1 := "true"
	b, e := strconv.ParseBool(str1) // 可以使用下划线 "_" 忽略error的返回：b, _ := strconv.ParseBool(str1)
	fmt.Printf("b=%t\n", b)
	fmt.Printf("e=%v\n", e)
	str2 := "12312414"
	num4, _ := strconv.ParseInt(str2, 10, 32)
	fmt.Printf("num4=%d, type=%T\n", num4, num4)
	str3 := "3.1415926"
	pai, _ := strconv.ParseFloat(str3, 64)
	fmt.Printf("pai=%e\n", pai)
}
