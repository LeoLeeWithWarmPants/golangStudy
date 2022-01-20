# **goroutine**

目前很多老牌语言多线程模型设计，都没有尽可能的最大化CPU的性能，结果就是多线程“多为并发”，而golang设计之初就为了尽可能的使用CPU的性能，尽量做到多核CPU并行执行提高执行效率。

当开启很多线程时，将会消耗大量的服务器资源，而golang可以轻轻松松开启上万个协程。

## **golang的主线程和协程**

go的主线程（实际上就是常规认为的程序线程）：

- 主线程是一个物理线程，直接作用在CPU上，是重量级的，线程的执行与切换将会消耗大量CPU资源。
- 可以启动多个协程，协程是轻量级的线程（编译器做的优化）。

go协程（轻量级的线程）：

- 从主线程开启，轻量级，是逻辑态，资源消耗小
- 有独立的栈空间
- 共享程序堆空间
- 调度由用户控制

## **MPG线程模型**

参考：[一文了解 Go 并发模型][1]

## **查看和设置CPU使用数**

```go
//获取当前环境的逻辑CPU核心数
cpuNum := runtime.NumCPU()
//GOMAXPROCS设置可同时执行的最大CPU数，并返回先前的设置。 若 n < 1，它就不会更改当前设置。本地机器的逻辑CPU数可通过 NumCPU 查询。本函数在调度程序优化后会去掉
runtime.GOMAXPROCS(actualProcsNum)
```

## **协程使用的线程安全问题**

协程的使用，必定会造成线程安全问题，比如多个协程同时操作同一个共享变量（全局变量），如果该变量是非线程安全的必将造成线程安全问题。如下：

```go
import (
	"fmt"
)

//该全局变量会存在线程安全的问题
var ResultMap = make(map[int]int, 10)

//计算一个数的阶乘放入map中
func calculateFactorial(num int) {
	result := 1
	for i := 1; i <= num; i++ {
		result *= i
	}
	ResultMap[num] = result
}

func main() {
	//计算1-200中，每一个数的阶乘，并将结果放入一个map
	for i := 1; i <= 200; i++ {
		//为每一个数字开启一个协程进行阶乘运算
		go calculateFactorial(i)
	}
    //输出map中的结果
	for idx, v := range ResultMap {
		fmt.Printf("ResultMap[%d]:%d\n", idx, v)
	}
}

```

结果：

```
fatal error: concurrent map writes

goroutine 33 [running]:
runtime.throw({0xeeaa8b, 0x0})
	C:/leolee/development/component/golang/go1.17.windows-amd64/go/src/runtime/panic.go:1198 +0x76 fp=0xc00013bf60 sp=0xc00013bf30 pc=0xe73416
runtime.mapassign_fast64(0x0, 0x0, 0x1b)
	C:/leolee/development/component/golang/go1.17.windows-amd64/go/src/runtime/map_fast64.go:101 +0x2c5 fp=0xc00013bf98 sp=0xc00013bf60 pc=0xe4f9c5
main.calculateFactorial(0x0)
	C:/leolee/development/workspace/personal/go/src/golangStudyProject/project25goroutine/demo2Thread-safe/main/main.go:17 +0x4d fp=0xc00013bfc8 sp=0xc00013bf98 pc=0xecdfed
main.main·dwrap·1()
	C:/leolee/development/workspace/personal/go/src/golangStudyProject/project25goroutine/demo2Thread-safe/main/main.go:25 +0x26 fp=0xc00013bfe0 sp=0xc00013bfc8 pc=0xece1c6
runtime.goexit()
	C:/leolee/development/component/golang/go1.17.windows-amd64/go/src/runtime/asm_amd64.s:1581 +0x1 fp=0xc00013bfe8 sp=0xc00013bfe0 pc=0xe9cf21
created by main.main
	C:/leolee/development/workspace/personal/go/src/golangStudyProject/project25goroutine/demo2Thread-safe/main/main.go:25 +0x5b
	
............
```

结果明显的指出，发生了并发的map写操作，然而golang的map并非是线程安全的。

也可以通过 **go build -race xxx.go** 对代码进行 “数据竞争检测”：

在go build -race执行后将会生成一个可执行文件（有可能被windows防火墙直接删除），运行该执行文件将会输出检测结果：

![线程不安全的数据竞争检测结果1](../static/img/线程不安全的数据竞争检测结果1.png)

发现了两处数据竞争（每次发生的数据竞争不一定一致，DDDD）

![线程不安全的数据竞争检测结果2](C:\leolee\development\workspace\personal\go\src\golangStudyProject\static\img\线程不安全的数据竞争检测结果2.png)

为了解决这样的问题，肯定是加锁，golang的sync包下提供了多线程同步的相关能力，其中可以使用 **sync.Mutex** 来对临界区进行加锁，类似于Java中的synchronized，属于重型锁，对于多线程高并发的访问，肯定有较大的影响。而且在demo2Thread-safe的示例中，还需要设置主线程的睡眠时间，来等待所有协程的执行完成，但是协程的执行时间很难确定。所以种种问题表明，使用 sync.Mutex 来对临界区加锁并不是一个好的办法来解决多协程并发访问共享变量的线程安全问题。

## **channel**

[一文了解 Go 并发模型][1] 一文中还提到了 **CSP并发模型**，golang中的 channel 就是 CSP并发模型 的具体实现。

- channel的本质就是一个数据结构（队列）
- 满足FIFO（先进先出原则）
- 本身是线程安全，多协程访问时，不需要加锁
- channel是有数据类型的（一个int类型的channel只能存取int类型的数据）

### **定义和声明**

var [变量名] chan [管道的数据类型]

- channel是引用类型
- 必须初始化才能写入数据，即使用前需要make，指定类型、size、capacity
- channel声明时指定的数据类型规定了该channel存取时的数据类型
- 放入管道的数据量最大为设置的容量，超过将会报错：deadlock!
- 在没有使用协程的情况下，如果我们的管道数据已经全部取出，继续获取将会报错：deadlock!









[1]: https://mp.weixin.qq.com/s?__biz=MzA5OTAyNzQ2OA==&amp;mid=2649709343&amp;idx=1&amp;sn=c4b51cf3d7dade9c799c3142dea6d06f&amp;chksm=8893647cbfe4ed6a83f4113b29b412121190cc42979e14d23b7db99af6d500c5a664c73c669c&amp;mpshare=1&amp;scene=1&amp;srcid=0711r5wHilGdt43uFZEVE5TD&amp;sharer_sharetime	"一文了解 Go 并发模型"

