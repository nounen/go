package main

import "fmt"

func main() {
	// defer 语句会延迟函数的执行直到上层函数返回
	// 延迟调用的参数会立刻生成，但是在上层函数返回前函数都不会被调用
	defer fmt.Println("world")

	fmt.Println("hello")
}
