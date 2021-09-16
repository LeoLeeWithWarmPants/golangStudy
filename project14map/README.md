# Map

与其他语言中的map类型基本相似，都是k-v形式的，k不可重复，元素是无序的，并且支持扩容。

基本语法

```go
var 变量名 map[keytype]valuetype
```

## key可以是很多类型

bool、数字、string、指针、channel，**还可以是只包含前面几种类型的 接口、结构体、数组**。

但是！！！**slice、map、function是不可以作为map的key，因为无法使用 == 判断key是否相等**。

## value的类型就没有特别的要求

## 声明

map的声明是不会分配内存的，初始化需要make，分配内存后才能赋值和使用。

