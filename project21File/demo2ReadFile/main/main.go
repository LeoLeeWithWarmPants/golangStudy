package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
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

//获取文件的bufferReader(默认是4096个字节的缓冲)，os.File实现了Reader接口
func getReader(file *os.File) *bufio.Reader {
	return bufio.NewReader(file)
}

func readFile(filePath string) {
	file := openFile(filePath)
	defer file.Close() //当函数执行完成后及时关闭file对象，防止内存泄漏

	reader := getReader(file)
	printFileByte(reader)
	fmt.Println("--------------------------------")
	printFileStr(reader) //?????不知道为何，重复读取同一个reader第二次就读取不到任何东西，哪怕是重新bufio.NewReader(file)产生新的reader也不行
	fmt.Println("--------------------------------")

}

func printFileStr(reader *bufio.Reader) {
	//循环读取reader中的内容
	for {
		strLine, err := reader.ReadString('\n') // 读取到一个换行符就结束
		if err != nil {
			//如果读取到文件的末尾，就跳出循环
			if err == io.EOF {
				break
			}
			fmt.Printf("reader read str err:%v\n", err)
		}
		fmt.Printf("%s", strLine)
	}
}

func printFileByte(reader *bufio.Reader) {
	var byteArr = make([]byte, 4, 8)
	for {
		byteNum, err := reader.Read(byteArr)
		if byteNum == 0 {
			break
		}
		if err != nil {
			fmt.Printf("reader read str err:%v\n", err)
		}
		fmt.Printf("%s", string(byteArr))
	}
}

func main() {
	readFile("C:\\Users\\WEIMOB\\Desktop\\fileTestForGolang.txt")

	//小文本可以直接全量读取
	content, err := ioutil.ReadFile("C:\\Users\\WEIMOB\\Desktop\\fileTestForGolang.txt")
	if err != nil {
		fmt.Printf("file read err:%v\n", err)
	}
	fmt.Printf("%v\n", string(content))
	fmt.Println("--------------------------------")
	fmt.Printf("%v\n", string(content))
}
