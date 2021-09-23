Golang中的方法比较特殊，其作用在指定的数据类型上，所以自定义类型（通过type定义的类型）都可以定义方法，不仅仅是struct，**如type interger int，可以给这个integer绑定方法**。

```go
type A struct {
	...
}
func (a A) xxx() {
	...
}
```

上述实例可解释为：

- func (a A) xxx()为A结构的一个方法（A结构有一个方法为func (a A) xxx()）。
- func (a A) xxx()与结构A是绑定的。

方法定义的语法：

```go
func ([recevier] [type]) [methodName]([params...]) ([return list...]) {
	//code
	return ...
}
```

注意：

- **一个对象调用自己的方法时，会将对象自身作为参数传递给方法**。
- **对象调用方法时，是将自身的副本作为参数进行传递的**。
  - **如果需要进行引用传递，那么就需要将recevier的类型定义为指针**。
- 方法的作用域控制规则与函数一致，方法名字母小写，只能在本包访问，方法名首字母大写，可以全局访问。
- **如果一个类型实现了一个String()这个方法，那么fmt.Println默认会调用这个类型的这个String()方法**。相当于Java重写了toString()

