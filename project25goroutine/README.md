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



[1]: https://mp.weixin.qq.com/s?__biz=MzA5OTAyNzQ2OA==&amp;mid=2649709343&amp;idx=1&amp;sn=c4b51cf3d7dade9c799c3142dea6d06f&amp;chksm=8893647cbfe4ed6a83f4113b29b412121190cc42979e14d23b7db99af6d500c5a664c73c669c&amp;mpshare=1&amp;scene=1&amp;srcid=0711r5wHilGdt43uFZEVE5TD&amp;sharer_sharetime	"一文了解 Go 并发模型"

