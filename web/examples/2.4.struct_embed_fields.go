package main

import "fmt"

type Human struct {
	name   string
	age    int
	weight int
}

// 说成 **继承** 更好理解
type Student struct {
	Human      // 匿名字段，那么默认Student就包含了Human的所有字段.
	speciality string
}

func main() {
	mark := Student{
		Human{
			"Lxl",
			25,
			120,
		},
		"Computer Science",
	}

	// 我们访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

	mark.name = "name 直接修改操作(不用通过 Human)"
	fmt.Println("His name is ", mark.name)

	mark.Human.age = -1
	fmt.Println("通过 Human 改 age, 但不是必须 ", mark.age)
}
