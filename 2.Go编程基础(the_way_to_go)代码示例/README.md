### 教程
* https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/directory.md

* http://wiki.jikexueyuan.com/project/the-way-to-go/


### 6.2.1 按值传递（call by value） 按引用传递（call by reference）
* 在函数调用时，像切片（slice）、字典（map）、接口（interface）、通道（channel）这样的 __引用类型__ 都是默认使用引用传递（即使没有显式的指出指针）。


### 6.2.2 命名的返回值（named return variables）
* eg: `func getX2AndX3_2(input int) (x2 int, x3 int) { ... }`


### 6.2.3 空白符（blank identifier）
* 空白符 `_` 用来匹配一些不需要的值，然后丢弃掉


### 6.3 传递变长参数
* 函数的最后一个参数是采用 `...type` 的形式，那么这个函数就可以处理一个变长的参数

* 如果参数被存储在一个 slice 类型的变量 slice 中，则可以通过 `slice...` 的形式来传递参数，调用变参函数。
