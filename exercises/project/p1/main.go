package main

import (
	"errors"
	"fmt"
	"golangStudyProject/exercises/project/p1/constant"
)

func start() {
	//键盘输入
	var key string = ""
label1:
	for {
		PrintMainMenu()
		fmt.Scanln(&key)
		isContinue := chooseMenu(key)
		if !isContinue {
			break label1
		}
	}
}

func chooseMenu(key string) bool {
	checkMenuKey(key)
	//是否继续执行的标识
	result := true
	switch key {
	case "1":
		PrintTitle(constant.MenuItemNames[0])
	case "2":
		PrintTitle(constant.MenuItemNames[1])
	case "3":
		PrintTitle(constant.MenuItemNames[2])
	case "4":
		PrintTitle(constant.MenuItemNames[3])
		result = false
	default:
		fmt.Println("请输入正确的菜单项序号")
	}
	return result
}

func checkMenuKey(key string) {
	if key == "" {
		errors.New("输入的菜单项为空，请输入菜单项序号...")
	}
}

func main() {
	start()
}
