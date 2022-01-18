package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	//该全局变量会存在线程安全的问题
	ResultMap = make(map[int]int, 10)
	//互斥锁，相当于Java中的synchronized
	lock sync.Mutex
)

//计算一个数的阶乘放入map中
func calculateFactorial(num int) {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}

	//加锁防止多线程操作map
	lock.Lock()
	//解锁
	defer lock.Unlock()
	ResultMap[num] = result
}

func main() {

	//计算1-50中，每一个数的阶乘，并将结果放入一个map
	for i := 1; i <= 50; i++ {
		//为每一个数字开启一个协程进行阶乘运算
		go calculateFactorial(i)
	}

	//由于不知道协程执行了多久，这里需要这只等待时间，防止在协程执行完成之前，主线程执行结束，退出程序
	time.Sleep(3 * time.Second)

	for idx, v := range ResultMap {
		fmt.Printf("ResultMap[%d]:%d\n", idx, v)
	}
}
