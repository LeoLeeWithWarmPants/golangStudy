# **类型断言**

所谓类型断言，就是再不知道接口变量具体类型的情况下，如果需要转换为具体类型时候，所使用的一种强制类型转换手段。其仍然是存在恐慌的风险（类似于运行时类型转换异常）。

可以采用可检测的类型断言方式，通过判断是否成功断言，来避免产生直接类型转换可能导致的恐慌发生。