package main

import "fmt"

type Usb interface {
	Start()
	Stop()
}

type Phone struct {
	name string
}

func (phone *Phone) Start() {
	fmt.Println("手机插入USB")
}

func (phone Phone) Stop() {
	fmt.Println("手机拔出USB")
}

//可以存在类型私有的方法
func (phone Phone) Call() {
	fmt.Println("拨打电话")
}

type Ipad struct {
	name string
}

func (ipad Ipad) Start() {
	fmt.Println("平板电脑插入USB")
}

func (ipad Ipad) Stop() {
	fmt.Println("平板电脑拔出USB")
}

//接口类型的数组体现多态的特点
func main() {
	//定义一个Usb接口类型的数组，存放Phone和Ipad的结构体变量
	var usbArray [2]Usb
	fmt.Printf("usbArray = %v\n", usbArray)
	usbArray[0] = &Phone{"Iphone13 Pro Max"} //由于Phone在实现start函数的时候参数设定为了指针，这里需要传递指针类型
	usbArray[1] = Ipad{"Ipad Pro"}
	fmt.Printf("usbArray[0] address : %p\n", &usbArray[0])
	fmt.Printf("usbArray[1] address : %p\n", &usbArray[1])
	fmt.Printf("usbArray = %v\n", usbArray)
	fmt.Printf("usbArray[0] : %v\n", usbArray[0])
	fmt.Printf("usbArray[1] : %v\n", usbArray[1])
}
