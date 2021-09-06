package main

import (
	"fmt"
	"time"
)

// MyError 自定义异常结构体
type MyError struct {
	option    string //异常操作
	errorTime string //异常时间
	msg       string //异常信息
}

func createErrorStruct(option string, msg string) *MyError {
	myErr := new(MyError)
	myErr.option = option
	myErr.msg = msg
	now := time.Now()
	myErr.errorTime = now.Format("2006-01-02 15:04:05")
	return myErr
}

func (myError *MyError) Error() string {
	errInfo := fmt.Sprintf("option:%s,msg:%s,time:%s\n", myError.option, myError.msg, myError.errorTime)
	return errInfo
}

// CheckAdult 检查年龄是否成年
func CheckAdult(age int) (bool, error) {
	if age >= 18 {
		fmt.Printf("年龄:%d，已经成年\n", age)
		return true, nil
	} else if age > 0 && age < 18 {
		fmt.Printf("年龄:%d，未成年\n", age)
		return false, nil
	} else if age == 0 {
		fmt.Println("输入的年龄为0，不合法")
		err := createErrorStruct("check age", "输入的年龄为0，不合法")
		return false, err
	} else {
		panic(createErrorStruct("check age", "输入的年龄小于0，不合法"))
	}
}
