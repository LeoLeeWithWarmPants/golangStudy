package main

import "fmt"

//二维数组声明
var tda1 [3][3]int

//二维数组的声明并初始化
var tda2 [2][2]int = [2][2]int{{1, 2}, {3, 4}}

//推断数组长度
var tda3 [2][2]int = [...][2]int{{1, 2}, {3, 4}}
var tda4 [2][...]int = [...][...]int{{1, 2}, {3, 4}}
var tda5 [3][...]int = [3][...]int{{1, 2}, {3, 4}}
var tda6 [3][...]int = [...][...]int{{1, 2}, {3, 4}, {}}

//推断类型
var tda7 = [...][...]int{{1, 2}, {3, 4}, {}}

func main() {
	fmt.Printf("与其他语言的二维数组一样，属于是一个数组的元素为数组类型,tda1=%v\n", tda1)
	//将左上角到右下角这条线上的数字赋值为1
	for i := 0; i < len(tda1); i++ {
		tda1[i][i] = 1
	}
	//输出二维数组
	for _, e := range tda1 {
		for idx, a := range e {
			if idx == len(e)-1 {
				fmt.Println(a)
			} else {
				fmt.Print(a, " ")
			}
		}
	}

	//二位数组的内存布局
	//二位数组中，每个元素的地址为一个一维数组的地址，且地址是连续的，这个地址同时是这个一维数组第一个元素的地址
	fmt.Printf("tda1[0] address=%p\n", &tda1[0]) // &tda1[0] = &(tda1[0])
	fmt.Printf("tda1[1] address=%p\n", &(tda1[1]))
	fmt.Printf("tda1[0][0] address=%p\n", &tda1[0][0])
	fmt.Printf("tda1[1][0] address=%p\n", &(tda1[1][0]))
}
