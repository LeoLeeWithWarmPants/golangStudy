package main

import (
	"errors"
	"fmt"
)

//1.errors.New("错误说明信息")：会返回一个error类型的值，标识一个错误异常
//2.panic内置函数，接收一个interface{}类型的值作为参数，可以接收error类型的变量，输出错误信息，并退出程序
func main() {
	test()
	fmt.Println("test()报错，当前Print也不会输出")
	fmt.Println("-----------------------------------------------")
	test2(-1) //参数大于0：执行正常；参数等于0：返回错误；参数小于0：panic终止执行
}

func checkBoolean(b bool) error {
	if b {
		fmt.Println("checkBoolean执行成功，没有异常")
		return nil
	}
	//返回自定义异常
	return errors.New("checkBoolean执行失败...")
}

// panic demo
func test() {
	err := checkBoolean(true) //输入false产生异常
	if err != nil {
		fmt.Println("checkBoolean执行异常")
		panic(err) //捕获异常，终止执行
	}
	fmt.Println("panic会退出程序的执行，所以当前Print并不会输出")
}

// 自定义错误demo
func test2(age int) {
	res, err := CheckAdult(age)
	if err == nil {
		fmt.Printf("年龄%d是否成年:%t\n", age, res)
	} else {
		fmt.Printf("err[%v]\n", err)
	}
}
