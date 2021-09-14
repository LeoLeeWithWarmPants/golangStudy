package main

import (
	"errors"
	"fmt"
)

const length int = 5

var intArray = [length]int{1, 3, 5, 7, 9}

//从一个数组中获取一个切片
//通过make创建的切片顶层的数组只能通过切片来操作，该数组没有其他的引用
func getSliceFromArray(array [length]int) []int {
	slice := array[1:3] // 1:3 代表从数组下标为1的元素开始取值，到下标为3的元素位置，不包含3。如果从0开始引用数组，也可以省略冒号前的0（如 array[:3]），反之同理。
	return slice
}

//通过make创建一个切片
func getSliceByMake() []int {
	var slice3 []int = make([]int, 4, 8)
	fmt.Printf("slice3 length=%d, capcity=%d\n", len(slice3), cap(slice3))
	return slice3
}

//创建切片时，初始化切片引用的数组
func getSliceByInitArray() []int {
	var slice4 []int = []int{1, 3, 5, 7, 9}
	fmt.Printf("slice4 length=%d, capcity=%d\n", len(slice4), cap(slice4))
	return slice4
}

//对切片进行切片，从一个切片产生一个新的切片
func getSliceFromSlice(slice []int) []int {
	return slice[0:2]
}

//切片的遍历
func sliceTraverseByFori(slice []int) {
	fmt.Println("sliceTraverseByFori executing...")
	for i := 0; i < len(slice); i++ {
		if i == len(slice)-1 {
			fmt.Printf("slice[%d]:%d \n", i, slice[i])
		} else {
			fmt.Printf("slice[%d]:%d ", i, slice[i])
		}
	}
	slice[0] = 2
}
func sliceTraverseByForRange(slice []int) {
	fmt.Println("sliceTraverseByForRange executing...")
	for idx, element := range slice {
		if idx == len(slice)-1 {
			fmt.Printf("slice[%d]:%d \n", idx, element)
		} else {
			fmt.Printf("slice[%d]:%d ", idx, element)
		}
	}
}

//测试切片扩容
func testDilatationByAppend(slice1 []int, slice2 []int, args ...int) []int {
	if slice1 == nil {
		panic(errors.New("被扩容的切片不能为空"))
	}
	if slice2 != nil {
		slice1 = append(slice1, slice2...)
	}
	if args != nil && len(args) > 0 {
		slice1 = append(slice1, args...)
	}
	return slice1
}

func main() {
	//从一个数组中获取一个切片
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

	//通过make创建切片
	slice3 := getSliceByMake()
	fmt.Printf("slice3:%v, slice3 memory address:%p\n", slice3, &slice3)

	//创建切片时，初始化切片引用的数组
	slice4 := getSliceByInitArray()
	fmt.Printf("slice4=%v\n", slice4)

	//切片fori遍历
	sliceTraverseByFori(slice4)
	sliceTraverseByFori(slice4) //由于切片是引用传递，所以方法内部对切片的变动，将会影响源数据
	sliceTraverseByForRange(slice4)

	//对切片进行切片，从一个切片产生一个新的切片
	fmt.Printf("slice4=%v\n", slice4)
	slice5 := getSliceFromSlice(slice4)
	fmt.Printf("slice5=%v\n", slice5)
	slice5[0] = 33
	fmt.Printf("after slice5 changed, slice4=%v\n", slice4) //说明切片slice4产生的切片slice5，slice4和slice5公用一个底层的数组

	//切片扩容测试
	intArray3 := [...]int{1, 2}
	slice6 := intArray3[:]
	fmt.Printf("slice6=%v, length=%d, capcity=%d, address=%p\n", slice6, len(slice6), cap(slice6), &slice6)
	slice7 := []int{3, 4, 5}
	slice6 = testDilatationByAppend(slice6, slice7)
	//由此可以看到切片在append元素之后，capcity增加了，虽然切片的地址没有变化，但是由数组声明后长度不可变可知，底层的数据肯定不是intArray3了
	fmt.Printf("slice6=%v, length=%d, capcity=%d, address=%p\n", slice6, len(slice6), cap(slice6), &slice6)

	//切片的拷贝
	slice8 := []int{1, 2, 3}
	slice9 := make([]int, 6)
	copyLength := copy(slice9, slice8)
	fmt.Printf("copyLength=%d\n", copyLength)
	fmt.Printf("slice8=%v\n", slice8)
	fmt.Printf("slice9=%v\n", slice9)
}
