package main

import (
	"fmt"
	u "golangStudyProject/project9func/demo1func/utils" //如果包名过长，可以设置包的别名，如别名u代替utils
)

//需要注意的是，在golang中，基本数据类型和数组都是值传递，即函数内改变形参不影响函数外的基本数据类型和数组的值
//golang中函数不支持重载
func main() {
	result := u.Calculate(float64(10), float64(2), u.Multiply)
	fmt.Printf("计算结果：%g\n", result)

	//多个返回值
	result1, result2 := u.Calculate2(float64(2), float64(4))
	fmt.Printf("result1=%g,result2=%g\n", result1, result2)
	//忽略返回值
	result3, _ := u.Calculate2(float64(4), float64(6))
	fmt.Printf("result3=%g\n", result3)

	//当形参是基本类型或者是数组时，想要函数内的改变函数外的变量值，可以传入变量的地址（&获取变量的地址），函数内以指针的方式操作该变量的值
	var num int8 = 1
	u.SelfIncrement(&num)
	fmt.Printf("after function [u.SelfIncrement] executed, variable [num] value is %d \n", num)

	//golang中，函数也是一种类型，函数可以赋值给一个变量，这个变量就是函数类型的变量，通过该函数变量可以实现对函数的调用
	selfIncrement := u.SelfIncrement
	selfIncrement(&num) //等价于u.SelfIncrement(&num)
	fmt.Printf("after function variable [selfIncrement] executed, variable [num] value is %d \n", num)

	//由于函数类型变量的存在，所以函数也可以作为一种形参进行传递
	res1 := u.TestFunctionFormalParameter(u.SelfIncrement, &num)
	fmt.Printf("res = %d\n", res1)

	//类型别名
	type myInt int
	var t1 myInt
	var t2 int
	t1 = 1
	t2 = int(t1) //这里仍然需要显式的类型转换，golang认为类型别名和原类型并不是一种类型
	fmt.Printf("t1=%d,t2=%d\n", t1, t2)

	//给函数类型参数定义类型别名，作为形参使用（TestFunctionFormalParameter2的形参）
	func1 := u.SelfIncrement
	/*type TestFuncType func(*int8)定义为全局变量，方便其他包调用*/
	res2 := u.TestFunctionFormalParameter2(func1, &num)
	fmt.Printf("res2=%d\n", res2)

	//golang支持对返回值命名，此时返回值列表必须加括号，return后面不需要指定返回的变量名
	res3 := u.TestReturnName()
	fmt.Printf("res3=%d\n", res3)

	//golang也支持可变参数
	res4 := u.Sum1(1, 2, 3)
	fmt.Printf("res4=%d\n", res4)
	//至少一个参数
	res5 := u.Sum2(1)
	fmt.Printf("res4=%d\n", res5)
}
