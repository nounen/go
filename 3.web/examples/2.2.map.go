package main

import (
	"fmt"
)

func main() {
	// 声明 map
	var map1 map[string]int
	// fmt.Println(map1)
	// TODO: 为什么不能直接开始进行 key value 赋值?
	map1 = map[string]int{"one": 1} // 感觉很繁琐
	map1["two"] = 2
	fmt.Println(map1)

	// map 的另一种声明方式
	map2 := make(map[string]int)
	// fmt.Println(map2)
	map2["one"] = 1
	map2["two"] = 2
	fmt.Println(map2)

	delete(map2, "one")
	fmt.Println(map2)
}
