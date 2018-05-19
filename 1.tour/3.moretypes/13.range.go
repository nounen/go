package main

import "fmt"

var pow = []int{1, 2, 4, 8, 16, 32, 64, 128}

func main()  {
	// for 循环的 range 格式可以对 slice 或者 map 进行迭代循环
	// 当使用 for 循环遍历一个 slice 时，每次迭代 range 将返回两个值
		// 第一个是当前下标（序号），第二个是该下标所对应元素的一个拷贝
	for i, v := range pow {
		fmt.Printf("2**%d = %d\n", i, v)
	}
}