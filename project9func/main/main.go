package main

import (
	"fmt"
	u "golangStudyProject/project8ProcessControl/utils" //如果包名过长，可以设置包的别名，如别名u代替utils
)

func main() {
	result := u.Calculate(float64(10), float64(2), u.Multiply)
	fmt.Printf("计算结果：%g\n", result)
}
