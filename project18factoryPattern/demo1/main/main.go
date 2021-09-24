package main

import (
	"fmt"
	"golangStudyProject/project18factoryPattern/demo1/model"
)

func main() {
	//通过工厂模式访问私有的struct
	p1 := model.GetPerson(int64(1), "Leo Lee", int8(26))
	fmt.Printf("p1 type:%T, p1 address:%p, p1 value:%v\n", p1, p1, *p1)
	p2 := model.GetPerson(int64(2), "Lydia", int8(28))
	fmt.Printf("p2 type:%T, p2 address:%p, p2 value:%v\n", p2, p2, *p2)

	//通过方法访问私有的字段
	p1.SetHeight(181)
	p1Height := p1.GetHeight()
	fmt.Printf("p1 height:%d\n", p1Height)
}
