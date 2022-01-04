package main

import (
	"fmt"
	"os"
)

func openFile(filePath string) *os.File {
	//打开文件
	file, err := os.Open(filePath)
	if err != nil {
		err.Error()
	}
	return file
}

func baseTest() {
	file := openFile("C:\\Users\\WEIMOB\\Desktop\\fileTestForGolang.txt")
	fmt.Printf("type:%T, value:%v\n", file, file)
	//关闭文件
	err := file.Close()
	if err != nil {
		fmt.Printf("close file err:%v\n", err)
	}
}

func main() {
	baseTest()
}
