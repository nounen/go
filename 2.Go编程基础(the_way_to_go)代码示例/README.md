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


### 6.4 defer 和追踪
* 关键字 defer 允许我们进行一些函数执行完成后的收尾工作，例如：
```go
// 1.关闭文件流 （详见 第 12.2 节）
// open a file  
defer file.Close()


// 2. 解锁一个加锁的资源 （详见 第 9.3 节）
mu.Lock()  
defer mu.Unlock() 


// 3.打印最终报告
printHeader()  
defer printFooter()


// 4.关闭数据库链接
// open a database connection  
defer disconnectFromDB()
```


### 6.5 内置函数
* `close`	    用于管道通信

* `len、cap`	len 用于返回某个类型的长度或数量（字符串、数组、切片、map 和管道）；cap 是容量的意思，用于返回某个类型的最大容量（只能用于切片和 map）

* `new、make`	new 和 make 均是用于分配内存：
    * `new` 用于值类型和用户定义的类型，如自定义结构，`make` 用于内置引用类型（切片、map 和管道）
    
    * 它们的用法就像是函数，但是将类型作为参数：new(type)、make(type)。
    
    * new(T) 分配类型 T 的零值并返回其地址，也就是指向类型 T 的指针（详见第 10.1 节）。它也可以被用于基本类型：v := new(int)。
    
    * make(T) 返回类型 T 的初始化之后的值，因此它比 new 进行更多的工作（详见第 7.2.3/4 节、第 8.1.1 节和第 14.2.1 节）new() 是一个函数，不要忘记它的括号

* `copy、append`	用于复制和连接切片

* `panic、recover`	两者均用于错误处理机制

* `print、println`	底层打印函数（详见第 4.2 节），在部署环境中建议使用 fmt 包

* `complex、real imag`	用于创建和操作复数（详见第 4.5.2.2 节）


### 接口
* 接口定义了一组方法（方法集），但是这些方法不包含（实现）代码：它们没有被实现（它们是抽象的）。接口里也不能包含变量。

* 类型不需要显式声明它实现了某个接口：接口被隐式地实现。多个类型可以实现同一个接口。

* 接口仅仅是定义方法，好像没啥用？
