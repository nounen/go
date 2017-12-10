package main

import (
	"fmt"
	"strconv"
)

type Element interface{}
type List []Element // Element 是空 interface, 所以 List slice 可以放任何类型的元素

type Person struct {
	name string
	age  int
}

//打印
func (p Person) String() string {
	return "(name: " + p.name + " - age: " + strconv.Itoa(p.age) + " years)"
}

func main() {
	list := make(List, 3)

	//  Element 是空 interface, 所以 List slice 可以放任何类型的元素
	list[0] = 1       //an int
	list[1] = "Hello" //a string
	list[2] = Person{"Dennis", 70}

	for index, element := range list {
		// value, ok = element.(T)，这里value就是变量的值，ok是一个bool类型，element是interface变量，T是断言的类型
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Println("list[%d] is of a different type", index)
		}
	}
}
