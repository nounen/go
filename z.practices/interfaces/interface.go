package main

import (
	"fmt"
	"math"
)

// TODO: go 没有 "implements" 关键字: 类型通过实现一个接口的所有方法来实现该接口
// TODO: 如果判断某个 类型 是否实现了某个接口?

// Abser 接口
type Abser interface {
	Abs() float64
}

// Vertex 结构体
type Vertex struct {
	X, Y float64
}

// Vertex 结构体实现接口的方法
func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("%T %v \n", v, v)
}
