package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("命令行参数：", len(os.Args), "个")
	//遍历os.Args切片就可以获取包括程序名之后的所有参数
	for idx, v := range os.Args {
		fmt.Printf("arg[%d]:%s\n", idx+1, v)
	}
}
