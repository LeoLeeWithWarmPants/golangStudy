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

}
