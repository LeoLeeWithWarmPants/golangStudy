package main

import "fmt"

//转义字符的在go中的使用
func main() {
	//制表符 \t
	fmt.Println("Hi Golang\tHello world")
	//换行 \n
	fmt.Println("aaa\nbbb")
	// \
	fmt.Println("aaa\\bbb")
	// "
	fmt.Println("aaa\"bbb")
	//回车，从当前行的最前面输出后面的字符，完全覆盖掉之前的内容
	fmt.Println("cccccc\rbbb")
}
