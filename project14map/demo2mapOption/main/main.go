package main

import "fmt"

func main() {
	//map删除元素，使用内置函数delete
	map1 := map[string]string{
		"a": "欸",
		"b": "币",
		"c": "西",
	}
	delete(map1, "a") // 若map为nil或无此元素，delete不进行操作
	delete(map1, "d")
	fmt.Printf("map1=%v, length=%d\n", map1, len(map1))
	//golang的map没有办法一次性删除所有的key-value，需要遍历删除
	for k, _ := range map1 {
		delete(map1, k)
	}
	fmt.Printf("map1=%v, length=%d\n", map1, len(map1))
	//或者对该map变量重新进行make初始化（指向新的内存地址），让该map变量对应的原来的引用对象成为垃圾。
	map1 = make(map[string]string)
	fmt.Printf("map1=%v, length=%d\n", map1, len(map1))

	//获取对应key的value
	map2 := map[string]string{
		"a": "欸",
		"b": "币",
		"c": "西",
	}
	val1, getRes1 := map2["a"]
	fmt.Printf("key:a, value:%s, getRes1:%v\n", val1, getRes1)
	val2, getRes2 := map2["d"]
	fmt.Printf("key:d, value:%s, getRes2:%v\n", val2, getRes2)

	//遍历，map的遍历只能使用for range方式
	for k, v := range map2 {
		fmt.Println(k, v)
	}
}
