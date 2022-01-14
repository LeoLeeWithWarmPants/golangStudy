Golang自带一个轻量级的测试框架testing，以及一个自带的go test命令来实现单元测试和性能测试。

## 创建单元测试：

用于单元测试的go文件，必须使用 _test.go 结尾，否则使用 go test 命令的时候将会提示找不到测试文件。测试文件中的单元测试方法也必须使用 Test 打头，否则也是无法识别为单元测试方法。

## 执行单元测试：

进入单元测试go文件的目录，在该目录下执行

```go
go test -v (-v 输出详情，包括成功的单元测试)
```

之后将会打印执行结果：

```
=== RUN   TestAutoIncrement
    main_test.go:13: 执行:bool
--- PASS: TestAutoIncrement (0.00s)
PASS
```

Goland中，对单元测试go文件所在文件夹或者更高的层级的文件夹右键，即可看到 go test 选项，将会执行该文件夹下所有的单元测试。

### go test 基本使用

通过 go help testflag 和  go help test 查看命令的使用说明

测试单个文件：

```shell
go test xxx_test.go xxx.go(该文件为被测试文件)
```

测试单个方法：

```
go test -test.run Testxxx(指定的单元测试方法名)
```

