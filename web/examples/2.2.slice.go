package main

import (
	"fmt"
)

func main()  {
	/*
	 * slice
	 */
	var slice1, slice2 []string
	// slice1[3] = "!" // slice 这么赋值不行! panic: runtime error: index out of range
	
	slice1 = append(slice1, "linl", "mou", "ren")
	
	fmt.Println(slice1)

	fmt.Println(len(slice1))
	
	fmt.Println(cap(slice1))
	
	// TODO: 为什么这里复制出来的 slice 是空的?
	var copyLen = copy(slice2, slice1)

	fmt.Println(slice2)

	fmt.Println(copyLen)
}