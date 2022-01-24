package main

import (
	"fmt"
	"golangStudyProject/project15OOP/demo1extends/models"
)

func main() {
	//声明管道，使用前必须make
	var intChan chan int = make(chan int, 2)
	// intchan=0xc0000d4180, intchan.address=0xc0000ce018, channel是一个引用类型，inchan只是存放了channel的地址
	fmt.Printf("intchan=%v, intchan.address=%p\n", intChan, &intChan)

	//声明只写channel，在默认情况下channel是双向的，可读可写
	var onlyWriteChan chan<- int
	onlyWriteChan = make(chan int, 2)
	onlyWriteChan <- 2
	//num := <-onlyWriteChan 报错：Invalid operation: <-onlyWriteChan (receive from the send-only type chan<- int)

	//声明只读channel,常用于参数声明，保证传入的channel会被正确的使用（读或者写）
	var onlyReadChan <-chan int
	onlyReadChan = make(chan int, 2)
	//onlyReadChan <- 2 报错：Invalid operation: onlyReadChan <- 2 (send to the receive-only type <-chan int)
	fmt.Println(onlyReadChan)

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

	//空接口类型的channel(泛型channel)
	genericChan := make(chan interface{}, 3)
	genericChan <- 1
	genericChan <- "LeoLee"
	//构建一个结构体，放入泛型channel
	person1 := models.Person{Name: "LeoLee", Age: 27}
	book1 := models.Book{Name: "golang从入门到放弃", Price: 100.02}
	person1.Book = book1
	person1.Price = 5000.02
	brand1 := &models.Brand{Name: "apple"}
	person1.Brand = brand1
	fmt.Printf("person1 = %v, phone's brand = %v\n", person1, (*person1.Brand).Name)
	genericChan <- person1
	//获取channel中的结构体
	<-genericChan //获取但是并不是赋值
	<-genericChan
	person2 := <-genericChan
	fmt.Printf("person2的类型为=%T, person2=%v\n", person2, person2)
	//能不能获取结构体的属性呢？？？
	//fmt.Printf("person2.Name=%v\n", person2.Name) //编译就报错，并不能直接通过person2获取，编译的时候认为这个person2就是interface{}类型
	//所以使用类型断言
	person3 := person2.(models.Person)
	fmt.Printf("person3.Name=%v\n", person3.Name)

	//channel的关闭
	strChan := make(chan string, 2)
	strChan <- "LeoLee"
	close(strChan)
	//关闭后只能读数据，不能写数据
	str1, ok := <-strChan
	fmt.Printf("str1=%v, ok=%v\n", str1, ok)
	//strChan <- "Lydia" //panic: send on closed channel

	//channel的遍历
	intChan2 := make(chan int, 10)
	for i := 0; i < 10; i++ {
		intChan2 <- i + 1
	}
	//for range遍历channel没有index参数，只有value,队列并不能通过下标去取值，只能按照FIFO的顺序取值
	//遍历时必须关闭channel，否则会出现fatal error: all goroutines are asleep - deadlock!
	close(intChan2)
	for v := range intChan2 {
		fmt.Println("v", v)
	}
}
