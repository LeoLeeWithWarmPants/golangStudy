package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	//func len(v Type) int，内建函数，在builtin包下，返回传入参数的长度，具体取决于参数的类型
	/*
		数组：v中元素的数量
		数组指针：*v中元素的数量（v为nil时panic）
		切片、映射：v中元素的数量；若v为nil，len(v)即为零
		字符串：v中字节的数量
		通道：通道缓存中队列（未读取）元素的数量；若v为 nil，len(v)即为零
	*/
	str := "hello LeoLee"
	fmt.Printf("len(str) = %d\n", len(str)) // len(str)=12

	//len输出带有中文的字符串长度
	//由于Go语言统一使用UTF-8编码，ASCII中字符（字母、数字、符号）占一个字节，中文占用3个字节，所以原本12个字节+三个汉字9个字节=21
	str += "哈哈哈"
	fmt.Printf("len(str + \"哈哈哈\") =%d\n", len(str)) // len(str + "哈哈哈") =21
	//golang中处理中文字符串
	strRune := []rune(str)
	for i := 0; i < len(strRune); i++ {
		if i < len(strRune)-1 {
			fmt.Printf("%c", strRune[i])
		} else {
			fmt.Printf("%c\n", strRune[i])
		}
	}

	//func Atoi(s string) (i int, err err)，字符串转数字
	i, err := strconv.Atoi("123a")
	if err == nil {
		fmt.Printf("i=%d\n", i)
	} else {
		fmt.Printf("err=%v\n", err)
	}

	//func Itoa(i int) string，整数转字符串
	intStr := strconv.Itoa(123)
	fmt.Printf("intStr=%s\n", intStr)

	//字符串转byte
	bytes := []byte("Hello哈哈哈")
	fmt.Printf("bytes=%v\n", bytes)

	//byte转字符串
	str2 := string([]byte{72, 101, 108, 108, 111, 229, 147, 136, 229, 147, 136, 229, 147, 136})
	fmt.Printf("str2=%s\n", str2)

	//十进制转其他进制 func FormatInt(i int64, base int) string
	v2 := strconv.FormatInt(int64(20), 2)
	fmt.Printf("二进制v2=%s\n", v2)
	v8 := strconv.FormatInt(int64(20), 8)
	fmt.Printf("八进制v8=%s\n", v8)
	v16 := strconv.FormatInt(int64(20), 16)
	fmt.Printf("十六进制v16=%s\n", v16)

	//查找字符串中是否包含另一个字符串
	var b = strings.Contains("123456", "234")
	fmt.Printf("123456 is contains 234:%t\n", b)

	//统计一个字符串在另一个字符串中一共出现的次数
	string1 := "e"
	string2 := "Hello LeoLee"
	count := strings.Count(string2, string1)
	fmt.Printf("count=%d\n", count)

	//不区分大小写比较字符串（==是区分大小写的）
	b2 := strings.EqualFold("ABC", "abC")
	fmt.Printf("b2=%t\n", b2)

	//返回某个字符串在另一个字符串中第一次出现的位置
	string3 := "eo"
	string4 := "Hello LeoLee"
	fmt.Printf("eo在%d的位置第一次出现\n", strings.Index(string4, string3))

	//返回某个字符串在另一个字符串中第一次出现的位置
	string5 := "eo"
	string6 := "Hello LeoLee LeoLee LeoLee"
	fmt.Printf("eo在%d的位置最后一次出现\n", strings.LastIndex(string6, string5))

	//字符串的替换
	string7 := "哈哈哈哈呵呵呵哈哈哈哈"
	string8 := "嘻嘻嘻"
	fmt.Println(strings.Replace(string7, "呵呵呵", string8, -1))

	//分割字符串
	string9 := "a,b,c,d,e,f,g"
	strArray := strings.Split(string9, ",")
	for i := 0; i < len(strArray); i++ {
		fmt.Println(strArray[i])
	}

	//大小写转换
	string10 := "Golang LeoLee"
	fmt.Println(strings.ToUpper(string10))
	fmt.Println(strings.ToLower(string10))

	//去除左右两边的空格
	string11 := "   哈哈哈哈   "
	fmt.Printf("去左右两边的空格:%s\n", strings.TrimSpace(string11))
	//去除左右两边指定的字符
	fmt.Printf("去除左右两边指定的字符:%s\n", strings.Trim("! hello ! ", " !"))

	//判断是否以指定的字符串结尾
	fmt.Println(strings.HasSuffix("xxx.jpg", "jpg"))
	//判断是否以指定的字符串开头
	fmt.Println(strings.HasPrefix("xxx.jpg", "xxx."))
}
