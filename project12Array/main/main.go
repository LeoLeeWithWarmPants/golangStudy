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

	//组数在golang中是基本数据类型，是值传递
	array6 := [...]int{1, 2, 3, 4}
	fmt.Printf("调用changeArray1前，array6=%v\n", array6)
	changeArray1(array6)
	fmt.Printf("调用changeArray1后，array6=%v\n", array6)
	//如果想要使用引用传递的方式传递数组变量，请使用指针
	fmt.Printf("调用changeArray2前，array6=%v\n", array6)
	changeArray2(&array6)
	fmt.Printf("调用changeArray2后，array6=%v\n", array6)
}

//由于组数在golang中是基本数据类型，是值传递，并非引用传递，所以该方法无法修改元数据array6的值
func changeArray1(array [4]int) {
	array[0] = 7
	fmt.Printf("changeArray1中array=%v\n", array)
}

//使用指针对数组进行引用传递
func changeArray2(array *[4]int) {
	array[0] = 8
	fmt.Printf("changeArray2中array=%v\n", array)
}
