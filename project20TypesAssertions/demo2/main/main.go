package main

import "fmt"

//通过类型断言判断变量的类型
func TypeOf(items ...interface{}) {
	for idx, e := range items {
		switch e.(type) {
		case bool:
			myPrint(idx, e)
		case float32:
			myPrint(idx, e)
		case float64:
			myPrint(idx, e)
		case int, int16, int32, int64:
			myPrint(idx, e)
		default:
			fmt.Println("暂不支持解析的数据类型")
		}
	}
}

func myPrint(i int, element interface{}) {
	fmt.Printf("第%d个参数的类型为：%T, value=%v\n", i+1, element, element)
}

func main() {
	param1 := 1
	param2 := 1.1
	param3 := true
	param4 := 1.234
	var param5 int64 = 30
	TypeOf(param1, param2, param3, param4, param5)
}
