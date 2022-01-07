package main

import (
	"bufio"
	"fmt"
	"os"
	"reflect"
	"strconv"
)

//----------判断文件或者文件夹是否存在----------

func checkFilePathExists(filePath string) bool {
	//通过获取指定路径的 info 来判断文件或者文件夹是否存在，如果返回*PathError，代表没有该文件或者文件夹
	fileInfo, err := os.Stat(filePath)
	fmt.Printf("path[%s] info: %v\n", filePath, fileInfo)
	if err != nil {
		if reflect.TypeOf(err).Kind() == reflect.TypeOf(new(os.PathError)).Kind() {
			fmt.Printf("[%s] is not exists\n", filePath)
		} else {
			fmt.Printf("check file path exists err: %v\n", err)
		}
		return false
	}
	return true
}

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

//清空文件
func clearFileContent(filePath string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_TRUNC, 777)
	if err != nil {
		fmt.Printf("openFile failed, err:%v\n", err)
		panic(err)
	}
	defer file.Close()
}

//追加数据
func appendContent(filePath string, appendContent string) {
	file, err := os.OpenFile(filePath, os.O_WRONLY|os.O_APPEND, 777)
	if err != nil {
		fmt.Printf("openFile failed, err:%v\n", err)
		panic(err)
	}
	defer file.Close()
	writer := getWriter(file)
	writer.WriteString(appendContent)
	writer.Flush()
}

func main() {
	createAndWriteFile("C:\\leolee\\development\\workspace\\personal\\go\\src\\golangStudyProject\\project21File\\demo3CreateAndWrite\\main\\test.txt", "LeoLee")
	// ./并不是代表当前文件所在路径的相对路径，还是代表了goPath的相对路径
	createAndWriteFile(".\\project21File\\demo3CreateAndWrite\\main\\test2.txt", "Lydia")

	appendContent("./project21File/demo3CreateAndWrite/main/test2.txt", "current line by append option")

	//clearFileContent("./project21File/demo3CreateAndWrite/main/test2.txt")

	fmt.Printf("is exists: %t\n", checkFilePathExists("./project21File/demo3CreateAndWrite/main/test2.txt"))
}
