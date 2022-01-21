package main

import "fmt"

func writeData(intDataChan chan int) {
	for i := 0; i < 30; i++ {
		fmt.Println("insert data:", i)
		intDataChan <- i
	}
	//写完数据就关闭管道（不影响管道的数据读取）
	close(intDataChan)
}

//读取数据（函数写法）
func readData(intDataChan chan int, exitChan chan bool) {
	for {
		v, ok := <-intDataChan
		fmt.Printf("%v\n", ok)
		if !ok {
			break
		}
		fmt.Printf("读取数据:%v\n", v)
	}
	//读取完成后修改退出channel的数据为true，告诉主线程可以继续执行
	exitChan <- true
	close(exitChan)
}

//gorountine和channel配合使用示例（多线程示例）
//一个协程读取管道数据，一个协程向管道内写入数据
func main() {
	//创建数据管道与退出管道
	intDataChan := make(chan int, 30)
	exitChan := make(chan bool, 1)

	//写入数据(匿名写法)
	go func() {
		for i := 0; i < 30; i++ {
			fmt.Println("insert data:", i)
			intDataChan <- i
		}
		//写完数据就关闭管道（不影响管道的数据读取）
		close(intDataChan)
	}()
	//go writeData(intDataChan)

	//读取数据(函数写法)
	go readData(intDataChan, exitChan)

	//主线程看从退出管道读取数据，管道内没有数据将会阻塞
	for {
		_, ok := <-exitChan
		if ok {
			break
		}
	}

	fmt.Println("主线程执行完毕")
}
