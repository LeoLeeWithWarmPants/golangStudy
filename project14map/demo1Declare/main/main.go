package main

import "fmt"

//声明一个map变量
var map1 map[string]string

func main() {
	//第一种方式：先声明，再make
	fmt.Printf("map1=%v\n", map1)
	fmt.Printf("map1 address=%p\n", &map1)     //变量map1是有内存地址的
	fmt.Printf("map1==nil? %v\n", map1 == nil) //在没有make之前map1==nil为true成立，代表没有分配内存空间
	//使用map要先make
	map1 = make(map[string]string, 2) //分配了2个空间
	map1["LeoLee"] = "18"
	fmt.Printf("map1=%v, length=%d\n", map1, len(map1))
	map1["Lydia"] = "18"
	fmt.Printf("map1=%v, length=%d\n", map1, len(map1))
	map1["LeoLee"] = "26"
	fmt.Printf("map1=%v, length=%d\n", map1, len(map1))
	map1["Tony"] = "1" //map会自动扩容
	fmt.Printf("map1=%v, length=%d\n", map1, len(map1))

	//第二种方式：声明map变量的同时，make
	map2 := make(map[string]int) //这里不定义初始size，根据内置函数make的api说明，会分配一个默认起始大小：0
	fmt.Printf("map2=%v, length=%d\n", map2, len(map2))
	map2["a"] = 1
	map2["b"] = 2
	fmt.Printf("map2=%v, length=%d\n", map2, len(map2))

	//第三种方式：声明直接赋值
	map3 := map[int]string{
		1: "c",
		2: "d",
	}
	fmt.Printf("map3=%v, length=%d\n", map3, len(map3))
}
