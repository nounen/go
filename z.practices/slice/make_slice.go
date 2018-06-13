package main

import "fmt"

// 切片可以用内建函数 make 来创建
func main() {
	a := make([]int, 5)
	// a len=5 cap=5 [0 0 0 0 0]
	printSlice("a", a)

	b := make([]int, 0, 5)
	// b len=0 cap=5 []
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}

func printSlice(s string, x []int) {
	fmt.Printf("%s len=%d cap=%d %v\n", s, len(x), cap(x), x)
}
