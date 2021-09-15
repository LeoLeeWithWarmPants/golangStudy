package main

import "fmt"

var myArray = [...]int{11, 34, 5, 90, 56, 1, 23}

const length = len(myArray)

//冒泡排序
func bubbleSort(array *[length]int) {
	var temp int
	for j := 0; j < length-1; j++ {
		for i := 0; i < length-1; i++ {
			if array[i] > array[i+1] {
				temp = array[i]
				(*array)[i] = (*array)[i+1]
				(*array)[i+1] = temp
			}
		}
	}
}

func main() {
	fmt.Printf("before sort, myArray=%v\n", myArray)
	bubbleSort(&myArray)
	fmt.Printf("after sort, myArray=%v\n", myArray)
}
