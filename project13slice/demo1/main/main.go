package main

import "fmt"

const length int = 5

var intArray = [length]int{1, 3, 5, 7, 9}

//从一个数组中获取一个切片
func getSliceFromArray(array [length]int) []int {
	slice := array[1:3] // 1:3 代表从数组下标为1的元素开始取值，到下标为3的元素位置，不包含3
	return slice
}

func main() {
	slice1 := getSliceFromArray(intArray)
	fmt.Printf("slice1 from an array, slice1=%v,length=%d,capacity=%d\n", slice1, len(slice1), cap(slice1))
	fmt.Printf("intArray memory address:%p\n", &intArray)
	fmt.Printf("intArray[1] memory address:%p\n", &intArray[1])
	fmt.Printf("slice1 memory address:%p\n", &slice1)
	fmt.Printf("slice1[0]=%d, slice1[0] memory address:%p\n", slice1[0], &slice1[0]) // slice1[0] == intArray[1] == 3

	//由于getSliceFromArray函数中声明并初始化了slice1，intArray是基本类型，通过值传递进入了getSliceFromArray函数，所以修改slice1的值并不影响intArray
	slice1[0] = 2
	fmt.Printf("after slice1 changed,slice1=%v\n", slice1)
	fmt.Printf("after slice1 changed,intArray=%v\n", intArray)
	//当slice2和intArray2处在同一作用域时，slice2中的元素地址指向了intArray2中的元素，修改slice2的元素值等于修改intArray2中的元素
	intArray2 := [length]int{0, 2, 4, 6, 8}
	slice2 := intArray2[1:3]
	fmt.Printf("slice2=%v\n", slice2)
	slice2[0] = 28
	fmt.Printf("after slice2 changed,slice2=%v\n", slice2)
	fmt.Printf("after slice2 changed,intArray2=%v\n", intArray2)
}
