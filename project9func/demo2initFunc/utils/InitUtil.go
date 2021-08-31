package utils

import "fmt"

var Name string
var Age int8

func init() {
	fmt.Println("InitUtil func init executing...")
	Name = "LeoLee"
	Age = 26
}
