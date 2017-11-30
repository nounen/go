package main

import "fmt"

/*
修改 map:
	在 map m 中插入或修改一个元素：
		m[key] = elem
		
	获得元素：
		elem = m[key]

	删除元素：
		delete(m, key)

	通过双赋值检测某个键存在： 
		如果 key 在 m 中， ok 为 true。否则， ok 为 false，并且 elem 是 map 的元素类型的零值。
		elem, ok = m[key]
*/
func main()  {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])
	
	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"]) // The value: 0

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok) // The value: 0 Present? false
}