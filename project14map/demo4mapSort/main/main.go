package main

import (
	"fmt"
	"sort"
)

//map排序
func main() {
	//在老版本中，使用print输出map也会是无序的，新版本中变为升序
	//但是遍历map仍然是无序的
	map1 := map[int]string{
		3: "b",
		2: "a",
		6: "c",
		5: "e",
	}
	fmt.Println("map1=", map1)
	for k, v := range map1 {
		fmt.Printf("%d:%s\n", k, v)
	}
	map2 := make(map[string]string)
	map2["b"] = "1"
	map2["a"] = "4"
	map2["g"] = "2"
	map2["f"] = "8"
	fmt.Println("map2=", map2)
	for k, v := range map2 {
		fmt.Printf("%s:%s\n", k, v)
	}

	//对map进行排序
	/*1.先将map的key放入切片中
	2.对切片进行排序
	3.遍历切片，按照切片遍历的顺序获取key对应的value*/
	var keys1 []int
	for k, _ := range map1 {
		keys1 = append(keys1, k)
	}
	fmt.Printf("keys1=%v\n", keys1)
	//使用sort包下的Ints函数，Ints函数将会把传入的int类型切片按照递增顺序排序
	sort.Ints(keys1)
	fmt.Printf("keys1=%v\n", keys1)
	for _, key := range keys1 {
		val, _ := map1[key]
		fmt.Println(key, ":", val)
	}
}
