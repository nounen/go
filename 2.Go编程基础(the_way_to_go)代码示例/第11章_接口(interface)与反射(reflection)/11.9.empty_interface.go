package main

import "fmt"

var i = 5
var str = "ABC"

type Person struct {
	name string
	age  int
}

type Any interface{}

// 空接口: 任何其他类型都实现了空接口
func main() {
	var val Any
	val = 5
	fmt.Printf("val has the value: %v\n", val)

	val = str
	fmt.Printf("val has the value: %v\n", val)

	pers1 := new(Person)
	pers1.name = "Rob Pike"
	pers1.age = 55

	val = pers1
	fmt.Printf("val has the value: %v\n", val)

	switch t := val.(type) {
	case int:
		fmt.Printf("Type int %T\n", t)
	case string:
		fmt.Printf("Type string %T\n", t)
	case bool:
		fmt.Printf("Type boolean %T\n", t)
	case *Person:
		fmt.Printf("Type pointer to Person %T\n", t)
	default:
		fmt.Printf("Unexpected type %T", t)
	}
}
