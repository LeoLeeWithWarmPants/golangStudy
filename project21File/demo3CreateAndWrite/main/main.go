package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

//----------创建文件并写入数据----------

func createFile(fileFullPath string) *os.File {
	//该方法最后的perm指定的是linux unix 等操作系统层面的权限，在windows系统中该参数无效，所有者权限_当前用户组权限_其他用户组的权限 r-4 w-2 x-1
	file, err := os.OpenFile(fileFullPath, os.O_WRONLY|os.O_CREATE|os.O_RDONLY, 666)
	if err != nil {
		fmt.Printf("openFile failed, err:%v\n", err)
		panic(err)
	}
	return file
}

func writeFile(file *os.File, content string) {
	writer := getWriter(file)
	for i := 0; i < 5; i++ {
		//bufio.Writer是带有缓冲区的，也就是写操作实际上是先将数据写入缓存
		writer.WriteString(strconv.FormatInt(int64(i+1), 10) + "." + content + "\n")
	}
	//将缓存中的数据真正写入文件
	writer.Flush()
}

func getWriter(file *os.File) *bufio.Writer {
	return bufio.NewWriter(file)
}

func createAndWriteFile(fileFullPath string, content string) {
	file := createFile(fileFullPath)
	defer file.Close()
	writeFile(file, content)
}

func main() {
	createAndWriteFile("C:\\leolee\\development\\workspace\\personal\\go\\src\\golangStudyProject\\project21File\\demo3CreateAndWrite\\main\\test.txt", "LeoLee")
	// ./并不是代表当前文件所在路径的相对路径，还是代表了goPath的相对路径
	createAndWriteFile(".\\project21File\\demo3CreateAndWrite\\main\\test2.txt", "Lydia")
}
