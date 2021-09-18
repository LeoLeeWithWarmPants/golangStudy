package main

import "fmt"

//坐标点
type coordinates struct {
	x, y int
}

//线
type line struct {
	startPoint, endPoint coordinates
	array                *[2]int
}

func main() {
	//定义一条直线
	l1 := &line{startPoint: coordinates{x: 1, y: 1}, endPoint: coordinates{x: 3, y: 6}}
	fmt.Printf("l1 type=%T\n", l1)
	fmt.Printf("l1=%v\n", l1)
	//通过内存地址可以看出这四个值在内存里是连续的：0xc0000101a0，0xc0000101a8，0xc0000101b0，0xc0000101b8
	//每个字段地址之间相差8个字节，int类型在64位操作系统中占用8个字节
	fmt.Printf("l1.startPoint.x=%v, address=%p\n", l1.startPoint.x, &l1.startPoint.x)
	fmt.Printf("l1.startPoint.y=%v, address=%p\n", l1.startPoint.y, &l1.startPoint.y)
	fmt.Printf("l1.endPoint.x=%v, address=%p\n", l1.endPoint.x, &l1.endPoint.x)
	fmt.Printf("l1.endPoint.y=%v, address=%p\n", l1.endPoint.y, &l1.endPoint.y)
	//如果结构体中的某个字段是指针类型，那么该指针对应的值就不一定是连续的了，但是指针自己的地址在该结构体中仍然是连续的
	//可以看到l1.array[0]和l1.array[1]的地址与l1.endPoint.y并不连续，但是l1.array与l1.endPoint.y连续
	a1 := [...]int{8, 9}
	(*l1).array = &a1
	fmt.Printf("l1.array address=%p\n", &l1.array)
	fmt.Printf("l1.array[0]=%v, address=%p\n", l1.array[0], &l1.array[0])
	fmt.Printf("l1.array[1]=%v, address=%p\n", l1.array[1], &l1.array[1])

}
