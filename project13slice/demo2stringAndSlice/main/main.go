package main

import "fmt"

//修改字符串，该方法无法处理中文字符（中文占三个字节）
func updateString(str string) string {
	byteArray := []byte(str)
	byteArray[0] = 'z'
	return string(byteArray)
}

//修改中文字符串
func updateChineseString(str string) string {
	runeArray := []rune(str)
	runeArray[0] = '李'
	return string(runeArray)
}

func main() {
	//string进行切片处理
	name := "Leo Lee"
	lastName := name[4:]
	fmt.Printf("lastName=%s\n", lastName)
	str := "哈哈哈嘻嘻嘻"
	xi := str[9:] //由于golang中的string类型中文字符占3个字节，所以这里是9开始
	fmt.Printf("xi=%s\n", xi)

	//修改字符串
	name2 := updateString(name)
	fmt.Printf("name2=%s\n", name2)

	name3 := "王老二"
	fmt.Printf("name3=%s\n", updateChineseString(name3))

}
