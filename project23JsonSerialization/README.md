由于Golang跨包必须使用大写的原因，所以要使用 json包中的Marshal函数，所有字段都必须大写（早前golang的版本中，不大写将会报错，较新版本中不大写将会被忽略）。

Golang中通过大小写来确定作用域，所以当一些变量序列化之后，其字段的名称的大小写也将会在json中体现出来 如下：

```json
json result:{"BaseInfo":{"Name":"LeoLee","Age":18},"Type":3,"Grade":"大一"}
json result:{"info":{"BaseInfo":{"Name":"LeoLee","Age":18},"Type":3,"Grade":"大一"},"name":"LeoLee"}
json result:[{"info":{"Name":"LeoLee","Age":18},"name":"LeoLee"},{"BaseInfo":{"Name":"LeoLee","Age":18},"Type":3,"Grade":"大一"}]
```

为了可以满足json中key的命名为首字母小写的驼峰命名法，我们可以使用statuct tag对结构体的字段进行标记，如：

```go
type Person struct {
	Name string `json:"name"`
	Age  int `json:"age"`
}

type Student struct {
	BaseInfo Person `json:"baseInfo"`
	Type     int // 1-小学生 2-中学生 3-大学生
	Grade    string
}
```

其中用

```
`json:""`
```

做标记的字段，序列化后将会使用tag指定的内容作为key。

