package main

import "fmt"

//Golang中的数组是值类型，并非是基本数据类型
var array [10]int
var array2 = [3]int{1, 2, 3}
var array3 = [...]int{2, 3, 4}        //让编译器推断数组的长度
var array4 = [4]int{0: 1, 1: 3, 2: 4} //指定下标进行赋值

func main() {
	//数组初始化之后，长度确定，每个元素的默认值为数组类型的默认值
	fmt.Printf("array=%v,length:%d\n", array, len(array))
	//数组的地址与数组的第一个元素的地址一致
	fmt.Printf("array的地址为%p\n", &array)
	fmt.Printf("array[0]的地址为%p\n", &array[0])
	//除数组第一个元素外的其他元素的地址=前一个元素的地址+数组类型所占用的字节数
	for i := 0; i < len(array); i++ {
		array[i] = i + 1
		fmt.Printf("array[%d]的地址=%p\n", i, &array[i])
	}
	fmt.Printf("array=%v,length:%d\n", array, len(array))
	fmt.Printf("array的地址为%p\n", &array)

	//数组的遍历
	array5 := [...]int{0: 1, 1: 5, 3: 4}
	for i := 0; i < len(array5); i++ {
		fmt.Printf("array5[%d]=%d\n", i, array5[i])
	}
	for idx, element := range array5 {
		fmt.Printf("array5[%d]=%d\n", idx, element)
	}

}
