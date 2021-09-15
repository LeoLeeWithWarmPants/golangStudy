string底层是一个byte数组，因此string也可以进行切片。

**由于string是不可变设计，不能直接通过 str[0] = 'a'的方式去修改字符串（编译报错）**。如果要修改string类型的变量，需要string to []byte/或者是[]rune，修改完成后再转成string。

