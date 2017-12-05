package main

import (
	"fmt"
)

// 函数支持多个返回值

func SumAndProduct(A, B int) (int, int) {
	return A + B, A * B
}

func main() {
	x := 3
	y := 4

	result1, result2 := SumAndProduct(x, y)

	fmt.Println(result1)
	fmt.Println(result2)
}