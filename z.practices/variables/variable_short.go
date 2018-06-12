package main

import (
	"fmt"
)

// 短变量声明： 只能用在函数内， 函数外使用 var
func main()  {
    i := 1

    fmt.Println(i)

    // %v 输出为相应值的默认格式
    fmt.Printf("%v \n", i)

    // %T 格式化输出该值的类型
    fmt.Printf("%T \n", i)
}