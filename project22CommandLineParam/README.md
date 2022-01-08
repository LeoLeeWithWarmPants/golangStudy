## **os.Args**

在os包中，有一个叫做 Args 的切片，用于获取go可执行文件通过命令执行时的命令行参数（**按照参数顺序获取，包含可执行程序名**）。

使用go build命令编译并源码为可执行文件后，在执行该文件后可自行添加命令行参数，用空格分隔：

```
.\TestGetCommandLineParam.exe a b c d
```

运行结果如下：

```
arg[1]:C:\leolee\development\workspace\personal\go\src\golangStudyProject\project22CommandLineParam\demo1\main\TestGetCommandLineParam.exe
arg[2]:a
arg[3]:b
arg[4]:c
arg[5]:d
```

这种方式的缺点就是没有参数名，命令行可读性较差。

## **flag.xxxVar**

使用flag包中的xxxVar等函数，可以指定参数名获取命令行参数，将获取到的参数值通过变量的指针进行赋值，需要注意一定要执行。

同时，支持指定命令参数的默认值

```go
flag.Parse()
```

运行结果如下：

```
PS C:\leolee\development\workspace\personal\go\src\golangStudyProject\project22CommandLineParam\demo2\main> .\TestGetCommandLineParam.exe
LeoLee 123456 127.0.0.1 3306
PS C:\leolee\development\workspace\personal\go\src\golangStudyProject\project22CommandLineParam\demo2\main> .\TestGetCommandLineParam.exe -u LeoLee2 -p 654321 -h 127.0.0.2 -port 6379
LeoLee2 654321 127.0.0.2 6379
PS C:\leolee\development\workspace\personal\go\src\golangStudyProject\project22CommandLineParam\demo2\main>

```

