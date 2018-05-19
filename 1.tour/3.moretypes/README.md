## 更多类型
* 指针, struct, slice, 映射


### 指针
* __指针__ 保存了 __值的内存地址__

* 类型 `*T` 是指向 `T` 类型值的指针。其零值为 `nil`

* 声明指针: `var p *int` // fmt.Println(p) 得到 `<nil>`

* __`&` 操作符__: 会生成一个指向其操作数的指针
    * 生成指针 (即值的内存地址)
    
    * eg:
    ```go
    i := 1
    fmt.Println(&i) // 得到 `0xc420016108` 指针保存了值的内存地址 
    ```

* __`*` 操作符__: 表示指针指向的底层值
    * 指向指针所在的值
    
    * eg:
    ```go
    i := 1
    fmt.Println(*&i) // 得到 `1`: 先取地址(先得到一个指针), 再指向地址所在的值
    ```

* Go _没有指针运算_


### 结构体
* 一个结构体（struct）就是 __一个字段__ 的集合

* 声明结构体: `type Xxx struct { ... }`

* 访问结构体字段: 结构体字段使用 _点号_ 来访问

* 结构体指针:
    * 结构体字段可以通过结构体指针来访问

    * 如果我们有一个指向结构体的指针 `p`，那么可以通过 `(*p).X` 来访问其字段 `X`

    * eg:
    ```go
    v := Vertex{1, 2}

	fmt.Println(v) // {1 2}
    fmt.Println(&v) // &{1 2}
    ```

* 结构体文法:
    * 赋值: 结构体文法通过直接列出字段的值来新分配一个结构体

    * 使用 Name: 语法可以仅列出部分字段。（字段名的顺序无关。）

    * 没有赋值的字段默认使用数据类型对应的 _零值_

    * __特殊__ 的前缀 `&` 返回一个指向结构体的指针(参考上面 `fmt.Println(&v) // &{1 2}`)


### 数组
* 类型 `[n]T` 表示拥有 `n` 个 `T` 类型的值的数组
    * 类型依旧放变量后面

* 数组的长度是其类型的一部分，因此 __数组不能改变大小__ (所以有了 _切片_ slice 类型)

* 数组批量赋值: `primes := [6]int{2, 3, 5, 7, 11, 13}`