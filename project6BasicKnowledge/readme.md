# 值类型和引用类型

- 值类型：基本数据类型（int、uint、float、bool、string）、数组以及结构体struct

  - 变量直接存储值，内存**通常**在栈中分配（也有例外的情况，姑且先这么理解）

- 引用类型：指针pointer、切片slice、map、管道chan、interface等都是引用类型

  - 变量存储的是一个地址，这个地址对应的内存空间才真正存储数据，内存**通常**在堆上分配，当没有任何变量引用这个地址时，该地址对应的数据空间便成为了一个垃圾，需要由GC回收。

  如下代码和图片为基本数据类型引用类型内存布局

  ```go
  //基本数据类型的内存布局
  var i8 int8 = 10
  //i的内存地址
  fmt.Printf("i8的内存地址:%v,&i8的类型:%T\n", &i8, &i8)
  var ptr *int8 = &i8
  fmt.Printf("ptr的值:%v,ptr的内存地址:%v,ptr的类型:%T\n", ptr, &ptr, ptr)
  //根据指针类型获取真正的值,使用*来获取
  fmt.Printf("real value:%d\n", *ptr)
  //通过指针来访问、修改指向的内存空间
  *ptr = 11
  fmt.Printf("i8=%d\n", i8)
  ```

![image-20210822212446151](https://github.com/LeoLeeWithWarmPants/golangStudy/blob/main/static/img/image-20210822212446151.png?raw=true)

# 标识符的基本使用

## 基本概念

Golang对各种变量、方法、函数等命名时使用的字符序列称为标识符

凡是可以自己命名的地方都可以叫标识符，如：var **num** int、var **str** string中的num和str。

## 标识符的命名规则

1. 由26个英文字母的大小写、0-9以及下划线“_”组成

2. 数字不可以开头

3. Golang中严格区分大小写，var num与var Num是声明了两个不同的变量

4. 标识符不能包含空格

5. 下划线“_”本身在Go中时一个特殊的标识符，可以代表任何其他的标识符，但是**它对应的值会被忽略，所以仅能被称为占位符使用，不能作为严格意义上的标识符使用**。

   如下：

   ```go
   str2 := "12312414"
   num4, _ := strconv.ParseInt(str2, 10, 32)
   ```

   下划线“_”作为占位符代表了error的返回值，只不过忽略了该返回值。

6. 不能以系统保留关键字作为标识符，如 break、if等关键字

   ![image-20210822215207613](https://github.com/LeoLeeWithWarmPants/golangStudy/blob/main/static/img/image-20210822215207613.png?raw=true)

   除了保留关键字之外，Golang还提供了36个预定标识符

   ![image-20210822220755855](https://github.com/LeoLeeWithWarmPants/golangStudy/blob/main/static/img/image-20210822220755855.png?raw=true)

## 标识符命名规范

1. 包名：保持package名称和当前.go文件所在目录名称一致，尽量采用简短、有意义的包名，不要和标准库包名冲突。
2. 使用驼峰命令法
3. Golang中没有访问权限关键字（Java中使用public、default、protected、private），如果变**量名、函数名、常量名首字母大写，那么就可以被其他包访问（public），如果首字母小写，则只能在本包中使用（private）**。

# Golang中的进制

- 二进制

  - 在Golang中，不能直接使用二进制来表示一个整数，它沿用了C的特点，需要使用 %b 来格式化

- 八进制

  ```go
  var x int8 = 012
  fmt.Printf("x=%d", x)
  //输出
  //x=10
  ```

- 十六进制

  ```go
  var y int8 = 0x11
  fmt.Printf("y=%d\n", y)
  //输出
  //y=17
  ```

# 运算符

与其他语言一样，Golang中也有如下类型的运算符

1. 算数运算符：+ - * / % ++ --等

   需要特别注意：

   ```go
   //算数运算符中需要特别注意除法，与Java一样，此处结果为2，省略了小数部分
   fmt.Println("10/4 = ", 10/4)
   //Java中float f = 10/4;f的结果是2.0，与Java一样，都不会是2.5，Golang中结果是2.000000
   var f float32 = 10 / 4
   fmt.Printf("10/4 = %f\n", f)
   //想要输出预期的小数，需要保证有浮点类型参与运算
   fmt.Printf("10/4 = %g\n", 10.0/4)
   fmt.Printf("10/4 = %g\n", 10/4.0)
   fmt.Printf("10/4 = %g\n", float32(10)/4)
   ```

   ```go
   //++ --在Golang中只能独立使用，否则编译报错
   var i8 int8 = 2
   //不能直接使用 a := i8++
   i8++
   a := i8
   fmt.Printf("a=%d\n", a)
   //在Golang中没有 ++i --i
   //甚至在if中都不可以使用
   /*if i8++ > 0 {
   	// do something
   }*/
   ```

2. 赋值运算符：+= -= *= /= %= 

   ```go
   //多重赋值
   var c = 20
   var b = 30
   c, b = b, c
   fmt.Printf("c=%d, b=%d\n",c ,b)
   ```

3. 比较运算符：== != > < >= <= <<= >>= &= ^= |=

4. 逻辑运算符：&& || !

5. 位运算：& | ^ << >>

6. **其他运算符：&（取指运算符）、*（取值运算符）**

   ```go
   var i8 int8 = 10
   //i的内存地址
   fmt.Printf("i8的内存地址:%v,&i8的类型:%T\n", &i8, &i8)
   var ptr *int8 = &i8
   fmt.Printf("ptr的值:%v,ptr的内存地址:%v,ptr的类型:%T\n", ptr, &ptr, ptr)
   //根据指针类型获取真正的值,使用*来获取
   fmt.Printf("real value:%d\n", *ptr)
   //通过指针来访问、修改指向的内存空间
   *ptr = 11
   fmt.Printf("i8=%d\n", i8)
   ```

7. **Golang没有三元运算符，需要使用if else代替**

