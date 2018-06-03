## 10 结构（struct）与方法（method）
* Go 通过 _类型别名_（alias types）和 _结构体_ 的形式支持用户自定义类型，或者叫定制类型

* 组成结构体类型的那些数据称为 __字段（fields）__

* 结构体的概念在软件工程上旧的术语叫 ADT（抽象数据类型：Abstract Data Type），在一些老的编程语言中叫 __记录（Record）__


## 10.1 结构体定义
* https://github.com/Unknwon/the-way-to-go_ZH_CN/blob/master/eBook/10.1.md

* 结构体定义的一般方式如下：
    ```go
    type identifier struct {
        field1 type1
        field2 type2
        ...
    }
    ```

* `type T struct {a, b int}` 也是合法的语法，它更适用于简单的结构体

* 使用 new
    *  使用 `new` 函数给一个新的结构体变量分配内存，它 _返回指向已分配内存的指针_：`var t *T = new(T)`

    * eg:
        ```go
        var t *T
        t = new(T) // 变量 t 是一个指向 T的指针，此时结构体字段的值是它们所属类型的零值
        ``

* 结构体的内存布局

* 递归结构体

* 结构体转换
