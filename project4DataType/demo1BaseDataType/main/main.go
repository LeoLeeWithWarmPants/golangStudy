package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//整数类型
	var i1 int8 = 1
	fmt.Printf("i1=%d\n", i1)
	var i2 int16 = 1
	fmt.Printf("i2=%d\n", i2)
	var i3 int32 = 1
	fmt.Printf("i3=%d\n", i3)
	var i4 int64 = 1
	fmt.Printf("i4=%d\n", i4)
	//uint8等是无符号的整数类型
	var ui uint8 = 2
	fmt.Printf("ui=%d\n", ui)
	a := 2 // 或者 var a = 2
	fmt.Printf("a data type:%T\n", a)
	b := 2 // 或者 var b = 2
	fmt.Printf("a need %d bytes\n", unsafe.Sizeof(b))

	//浮点类型（golang的浮点类型同样会丢精度）
	var f32 float32 = 123.456
	fmt.Printf("f32=%G\n", f32)
	var f64 float32 = 12398.12312
	fmt.Printf("f64=%G\n", f64)
	//隐式的声明浮点类型，默认是float64
	n := 3.1415926 // var n = 3.1415926
	fmt.Printf("n data type:%T\n", n)

	//字符类型
	var char1 byte = 'a'
	char2 := 'b'
	fmt.Printf("char1=%v,char2=%v\n", char1, char2) // %v：值的默认格式表示，这里输出:char1=97,char2=98
	fmt.Printf("char1=%c,char2=%c\n", char1, char2) // %c：值对应的unicode码值，这里输出：char1=a,char2=b
	var char3 int16 = '李'
	fmt.Printf("char3=%v\n", char3)
	//byte 0~255，一个字节，无法存放汉字，使用int代替byte
	fmt.Printf("char3=%c\n", char3)
	//字符类型可以进行运算，相当于整数参与运算，应为字符都存在其对应的Unicode码
	fmt.Printf("1 + char3 = %d\n", 1+char3)

	//布尔类型
	bbb := false
	fmt.Printf("b=%t\n", bbb)                              // %t：true或false
	fmt.Printf("bool need %d bytes\n", unsafe.Sizeof(bbb)) // bool占用一个字节

	//字符串类型,go中的字符串是不可变的
	var city string = "上海"
	fmt.Printf("city:%s\n", city)
	//字符串的拼接
	name := "Leo Lee"
	fmt.Printf("%s\n", city+name)
	area := "浦东新区"
	address := city + area + "陆家嘴"
	fmt.Printf("address:%s\n", address)
	//当拼接的字符串过长时，可以换行，但是+号要留在上一行，否则语法报错
	str := "a" + "b" + "c" +
		"d" + "e"
	fmt.Printf("%s\n", str)
}
