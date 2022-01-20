package main

import "fmt"

func main() {
	//声明管道，使用前必须make
	var intChan chan int = make(chan int, 2)
	// intchan=0xc0000d4180, intchan.address=0xc0000ce018, channel是一个引用类型，inchan只是存放了channel的地址
	fmt.Printf("intchan=%v, intchan.address=%p\n", intChan, &intChan)

	//向管道写入数据(写入的数据量最大不能超过初始化时候指定的容量)
	intChan <- 1 //直接写入数据
	num1 := 2
	intChan <- num1
	//intChan <- 3 //fatal error: all goroutines are asleep - deadlock!

	//管道的长度和容量
	fmt.Printf("intChan len=%d, cap=%d\n", len(intChan), cap(intChan))

	//从管道获取数据(FIFO，当channel中没有数据的时候，再次获取数据将会报错)
	num2 := <-intChan
	fmt.Printf("num2=%d\n", num2)
	num3 := <-intChan
	fmt.Printf("num3=%d\n", num3)
	//num4 := <-intChan //fatal error: all goroutines are asleep - deadlock!
}
