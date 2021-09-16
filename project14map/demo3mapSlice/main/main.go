package main

import "fmt"

//map切片，类似于Java中List<Map>
func main() {
	personSlice := make([]map[string]string, 2, 2)
	personSlice[0] = make(map[string]string)
	personSlice[0]["name"] = "LeoLee"
	personSlice[0]["age"] = "26"
	personSlice[1] = make(map[string]string)
	personSlice[1]["name"] = "Lydia"
	personSlice[1]["age"] = "29"
	//如果继续追加，会有越界
	/*personSlice[2] = make(map[string]string)
	personSlice[2]["name"] = "Lydia"
	personSlice[2]["age"] = "29"*/
	//所以需要使用slice的append，这样切片可以动态扩容
	personMap1 := make(map[string]string)
	personMap1["name"] = "Tony"
	personMap1["age"] = "30"
	personSlice = append(personSlice, personMap1)
	for _, m := range personSlice {
		fmt.Println(m)
	}
}
