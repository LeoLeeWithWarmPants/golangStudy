init函数总是在main函数之前执行，通常在init函数中执行一些初始化操作。

```go
package main
import "fmt"
func init() {
   fmt.Println("func init executing...")
}
func main() {
   fmt.Println("func main executing...")
}
```

如果一个文件同时包含了全局变量、init函数、main函数，那么执行的顺序是：

1. 全局变量的定义
2. init函数
3. main函数

```go
package main

import "fmt"

//全局变量定义首先初始化，其次执行init函数，最后是main函数
var num = initNum()

func initNum() int {
	fmt.Println("func initNum executing...")
	return 20
}

//init函数总是在main函数之前执行，通常在init函数中进行一些初始化操作
func init() {
	fmt.Println("func init executing...")
	fmt.Printf("num=%d\n", num)
}

func main() {
	fmt.Println("func main executing...")
	fmt.Printf("num=%d\n", num)
}
```

每一个go文件中的init函数在该文件被引用的时候都会执行，如