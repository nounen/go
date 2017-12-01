package main

import (
	"fmt"
	"math"
)

type MyFloat float64

// 你可以对包中的 任意 类型定义任意方法，而不仅仅是针对结构体
func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}

	return float64(f)
}

func main()  {
	// Sqrt2 是 math 包里的一个常量
	f := MyFloat(-math.Sqrt2)

	fmt.Println(f.Abs())
}