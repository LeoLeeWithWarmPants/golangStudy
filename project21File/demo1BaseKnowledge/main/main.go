package main

import (
	"fmt"
	"os"
)

func openFile(filePath string) *os.File {
	file, err := os.Open(filePath)
	if err != nil {
		err.Error()
	}
	return file
}

func main() {
	file := openFile("C:\\Users\\WEIMOB\\Desktop\\fileTestForGolang.txt")
	fmt.Printf("type:%T, value:%v\n", file, file)
}
