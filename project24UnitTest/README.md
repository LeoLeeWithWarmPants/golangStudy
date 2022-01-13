Golang自带一个轻量级的测试框架testing，以及一个自带的go test命令来实现单元测试和性能测试。

## 创建单元测试：

用于单元测试的go文件，必须使用 _test.go 结尾，否则使用 go test 命令的时候将会提示找不到测试文件。测试文件中的单元测试方法也必须使用 Test 打头，否则也是无法识别为单元测试方法。

## 执行单元测试：

进入单元测试go文件的目录，在该目录下执行

```go
go test -v
```

之后将会打印执行结果：

```
=== RUN   TestAutoIncrement
    main_test.go:13: 执行:bool
--- PASS: TestAutoIncrement (0.00s)
PASS
```

Goland中，对单元测试go文件所在文件夹或者更高的层级的文件夹右键，即可看到 go test 选项，将会执行该文件夹下所有的单元测试。

