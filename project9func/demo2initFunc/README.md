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

**每一个go文件中的init函数在该文件被引用的时候都会执行**，如[这个示例](https://github.com/LeoLeeWithWarmPants/golangStudy/blob/main/project9func/demo2initFunc/main/main.go)中main函数中引用了其他包中的全局变量，这些变量被自己所在go文件中的init函数初始化。

执行顺序为：

1. 按照引用链，首先执行被引用go文件中的全局变量定义，其次执行该文件中的init函数
2. 之后执行引用文件的中全局变量定义，之后执行该文件中的init函数，如果当前文件仍然被其他文件应用，按照引用链继续执行。
3. 直到main函数所在文件中，仍然是先执行全局变量的初始化，之后是当前文件的init函数。

