package main

import (
	"errors"
	"fmt"
)

//获取1-10000之间的所有素数

func checkPrimeNum(num int) (bool, error) {
	if num == 1 {
		return true, nil
	}
	if num < 1 {
		return false, errors.New("素数是大于1的自然数")
	}

	var flag bool = true
	for i := 2; i < num; i++ {
		if num%i == 0 {
			flag = false
			break
		}
	}
	return flag, nil
}

func checkFinish(exitChan chan bool, goroutineNum int, primeNumChan chan int) {
	for i := 0; i < goroutineNum; i++ {
		//channel取不到值的时候将会阻塞
		<-exitChan
	}
	close(exitChan)
	fmt.Printf("所有协程均已完成，共获取%v个素数\n", len(primeNumChan))
}

func putNum(intChan chan int) {
	for i := 1; i <= 8000; i++ {
		intChan <- i
	}
	//关闭管道
	close(intChan)
}

func primeNumFilt(intChan chan int, primeNumChan chan int, exitChan chan bool, taskName string) {
	var processNum int
	for {
		num, ok := <-intChan
		if !ok {
			//如果取不到数了就退出
			break
		}
		isPrimeNum, err := checkPrimeNum(num)
		if err != nil {
			fmt.Printf("checkPrimeNum error, msg:%v\n", err)
			panic(err)
		}
		if isPrimeNum {
			//将素数放入primeChan
			primeNumChan <- num
			processNum++
		}
	}
	//向退出的管道写入标志
	exitChan <- true
	fmt.Printf("任务协程%s执行完成，找到了%d个素数\n", taskName, processNum)
}

func printPrimeNum(primeNumChan chan int) {
	for v := range primeNumChan {
		fmt.Println(v)
	}
}

func main() {
	//待筛选管道
	intChan := make(chan int, 1000)
	//筛选结果管道
	primeNumChan := make(chan int, 3000)
	//退出标识
	exitChan := make(chan bool, 4)
	//开启一个协程放入素数
	go putNum(intChan)
	//开启四个协程获取并筛选素数
	var goroutineNum int = 4
	for i := 0; i < goroutineNum; i++ {
		go primeNumFilt(intChan, primeNumChan, exitChan, "任务"+fmt.Sprintf("%d", i+1))
	}
	checkFinish(exitChan, goroutineNum, primeNumChan)
	//关闭channel
	close(primeNumChan)
	printPrimeNum(primeNumChan)
}
