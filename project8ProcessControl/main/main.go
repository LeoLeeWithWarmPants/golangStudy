package main

import (
	"fmt"
	"math/rand"
	"time"
)

//流程控制
func main() {
	//if-else
	var age int8
	age = 18
	//if的括号可以省略
	if age > 18 {
		fmt.Println("成年人")
	} else if age == 18 {
		fmt.Println("刚刚成年")
	} else {
		fmt.Println("未成年")
	}
	//流程语句中声明变量并赋值
	if age2 := 20; age2 > 18 {
		fmt.Println("成年人")
	}

	//switch，golang的swtich中，case后的表达式可以是多个，用逗号分开；case语句块不需要写break，默认会有break
	week := 5
	switch week + 1 {
	case 1:
		fmt.Println("一")
	case 1 - 1:
		fmt.Println("零")
	case 2:
		fmt.Println("二")
	case 3:
		fmt.Println("三")
	case 4, 5:
		fmt.Println("四或者五")
	default:
		fmt.Println("？？？")
	}
	//switch后可以不写表达式，当作if-else使用
	a := 1
	b := 18
	switch {
	case a == 1:
		fmt.Println("aaa")
	case b == 18:
		fmt.Printf("bbb")
	default:
		fmt.Println("相当于else")
	}
	//switch后可以直接声明一个变量，分号结束(这么做没有什么意义啊)
	switch c := 1; {
	case c == 1:
		fmt.Println("cccc")
	default:
		fmt.Println("这是什么鬼语法")
	}
	//fallthrough穿透，当前case执行完成后，强制执行下一个case中的代码块，甚至穿透到default
	d := 1
	switch d {
	case 1:
		fmt.Println(1)
		fallthrough
	case 2:
		fmt.Println(2)
		fallthrough
	case 3:
		fmt.Println(3)
		fallthrough
	default:
		fmt.Println("fallthrough穿透")
	}
	//注意事项
	switch 'a' {
	case 'a':
		fmt.Println("编译通过")
	}
	var bbb byte = 'a'
	switch get(bbb) + 1 {
	case 'a':
		fmt.Println("编译通过")
	}

	//for
	for i := 0; i < 10; i++ {
		fmt.Printf("第%d次循环\n", i+1)
	}
	i := 1
	for i <= 10 {
		fmt.Printf("第%d次循环\n", i)
		i++
	}
	//死循环，等价于 for ;; {}，通常需要break跳出
	/*for ;; {

	}
	for {

	}*/
	//break
	count := 1
	var num int
	for {
		//设置随机种子
		rand.Seed(time.Now().UnixNano())
		num = rand.Intn(100) + 1
		if num == 99 {
			fmt.Println("随机到了99")
			break
		}
		count++
	}
	//嵌套for循环时，break只跳出当前for循环，可以指定break跳出指定label对应的for循环
label1:
	for i := 0; i < 5; i++ {
		for j := 0; j < 10; j++ {
			if j == 3 {
				break label1
			}
			fmt.Println("j=", j)
		}
	}
	//continue用于结束本次循环，下一次循环还会继续执行
	for i := 0; i < 10; i++ {
		if i == 4 {
			fmt.Println("i=4的循环结束")
			continue
		}
		fmt.Println("i=", i)
	}
	//continue结束本次循环，从指定here标签标记出开始下一次循环
here1:
	for i := 0; i < 2; i++ {
		for j := 0; j < 4; j++ {
			if j == 2 {
				continue here1
			}
			fmt.Println("i=%d,j=%d", i, j)
		}
	}
	//遍历数组或者是集合
	str := "Hello world"
	for i := 0; i < len(str); i++ {
		fmt.Printf("str[%d]=%q\n", i, str[i])
	}
	//for range，增强型for循环，与java中类似
	for idx, value := range str {
		fmt.Printf("str[%d]=%q\n", idx, value)
	}

	//Golang的goto可以无条件的将执行位置转移到指定行
	//goto通常与条件语句配合使用，可用来实现条件转移，跳出循环体等功能
	condition := 1
	fmt.Println("step1")
	if condition == 1 {
		goto point1
	}
	fmt.Println("step2")
point2:
	fmt.Println("step3")
	condition = 4
point1:
	fmt.Println("step4")
	if condition != 4 {
		condition = 3
	}
	fmt.Println("step5")
	if condition == 3 {
		goto point2
	}
	fmt.Println("step6")
}
func get(b byte) byte {
	return b + 1
}
