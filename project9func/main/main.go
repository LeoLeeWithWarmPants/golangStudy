package main

import (
	"fmt"
	u "golangStudyProject/project9func/utils" //如果包名过长，可以设置包的别名，如别名u代替utils
)

func main() {
	result := u.Calculate(float64(10), float64(2), u.Multiply)
	fmt.Printf("计算结果：%g\n", result)

	//多个返回值
	result1, result2 := u.Caculate(float64(2), float64(4))
	fmt.Printf("result1=%g,result2=%g\n", result1, result2)
	//忽略返回值
	result3, _ := u.Caculate(float64(4), float64(6))
	fmt.Printf("result3=%g\n", result3)
}
