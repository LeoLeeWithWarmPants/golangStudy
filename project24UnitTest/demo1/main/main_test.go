package main

import (
	"testing" //引入golang的testing组件
)

func TestAutoIncrement(test *testing.T) {
	resultInt := AutoIncrement(1)
	if resultInt != 2 {
		test.Fatalf("AutoIncrement(1) 执行错误，预期返回结果：%d，实际返回结果:%d\n", 2, resultInt)
	}
	//如果执行符合预期
	test.Logf("执行:%T\n", resultInt == 1)
}
