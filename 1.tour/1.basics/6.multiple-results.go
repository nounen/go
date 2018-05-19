package main

import "fmt"

// 多个返回值, 每个返回值的类型都要声明
func swap(x, y string) (string, string) {
	return y, x
}

func main()  {
	// := 是专门接收多个返回值的符号?
	a, b := swap("hello", "world")

	fmt.Println(a, b)
}