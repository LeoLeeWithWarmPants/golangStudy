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

