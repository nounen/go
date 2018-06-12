package main

import "fmt"

// Vertex 坐标
type Vertex struct {
	X int
	Y int
}

func main() {
	v := Vertex{1, 2}

	p := &v

	// TODO: 这里明明只是取地址, 怎么就能直接赋值了, 应该是 *p.X = 99 来赋值才对
	// __不过这么写太啰嗦了，所以语言也允许我们使用 __隐式间接引用__， 直接写 p.X 就可以__
	p.X = 99

	fmt.Println(v)
}
