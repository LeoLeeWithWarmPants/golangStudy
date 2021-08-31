package utils

import "fmt"

// Add 变量名大写，可以被其他包中代码使用
var Add byte = '+'
var Subtract byte = '-'
var Multiply byte = '*'
var Divide byte = '/'

// TestFuncType 给函数类型参数定义别名
type TestFuncType func(*int8)

// Calculate 根据 param option对num1和num2进行运算（函数名第一个字母大写，将可以跨包被调用，否则只能在当前包中使用）
func Calculate(num1 float64, num2 float64, option byte) float64 {
	var result float64
	switch option {
	case '+':
		result = num1 + num2
	case '-':
		result = num1 - num2
	case '*':
		result = num1 * num2
	case '/':
		result = num1 / num2
	default:
		fmt.Println("unsupported option")
	}
	return result
}

// Calculate2 多个返回值，如果想忽略某个返回值可以使用_作为占位符
func Calculate2(num1 float64, num2 float64) (float64, float64) {
	return num1 + num2, num1 * num2
}

// SelfIncrement 将目标数自增1
func SelfIncrement(num *int8) {
	*num++
}

// TestFunctionFormalParameter 函数类型变量作为形参
func TestFunctionFormalParameter(funcParameter func(*int8), num *int8) int8 {
	funcParameter(num)
	return *num + 10
}

func TestFunctionFormalParameter2(funcType TestFuncType, num *int8) int8 {
	funcType(num)
	return *num + 10
}

// TestReturnName golang支持对返回值命名，此时返回值列表必须加括号，return后面不需要指定返回的变量名
func TestReturnName() (rn int8) {
	rn = 2
	return
}

// Sum1 golang也支持可变参数
func Sum1(args ...int) int {
	if len(args) > 0 {
		result := 0
		for _, num := range args {
			result += num
		}
		return result
	}
	return 0
}

// Sum2 至少一个参数
func Sum2(num1 int, nums ...int) int {
	result := 0
	for _, num := range nums {
		result += num
	}
	return result + num1
}
