package main

import "fmt"

type Human struct {
	name string
	age int
	phone string
}

type Student struct {
	Human // 匿名字段 (继承)
	school string
}

type Employee struct {
	Human
	company string
}

// 这里 (h Human) 或者 (h *Human) 都是一样的 (所以没事别乱用指针类型, 看着还是有点蛋疼)
func (h Human) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

func main() {
	mark := Student{Human{"Mark", 25, "222-222-YYYY"}, "MIT"}
	sam := Employee{Human{"Sam", 45, "111-888-XXXX"}, "Golang Inc"}

	mark.SayHi()
	sam.SayHi()
}