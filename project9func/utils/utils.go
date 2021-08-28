package utils

import "fmt"

// Add 变量名大写，可以被其他包中代码使用
var Add byte = '+'
var Subtract byte = '-'
var Multiply byte = '*'
var Divide byte = '/'

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

// Caculate 多个返回值，如果想忽略某个返回值可以使用_作为占位符
func Caculate(num1 float64, num2 float64) (float64, float64) {
	return num1 + num2, num1 * num2
}
