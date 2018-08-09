package main

import (
	"math/rand"
	"fmt"
)

type tree struct {
	value int
	left, right *tree
}

// 使用一个二叉树来实现一个插入排序
// TODO: 看不懂
func main() {
	data := make([]int, 10)

	// 初始化
	fmt.Println(data)

	for i := range data {
		data[i] = rand.Int() % 50
	}

	// 源数据
	fmt.Println(data)

	Sort(data)
	// 排序后
	fmt.Println(data)

}

// 使用一个二叉树来实现一个插入排序
func Sort(values []int)  {
	var root *tree

	for _, v := range values{
		root = add(root, v)
	}

	appendValues(values[:0], root)
}

// TODO: 看不懂
func appendValues(values []int, t *tree) []int {
	if t != nil {
		values = appendValues(values, t.left)
		values = append(values, t.value)
		values = appendValues(values, t.right)
	}

	return values
}

// TODO: 看不懂
func add(t *tree, value int) *tree {
	if t == nil {
		t = new(tree)
		t.value = value
		return t
	}

	if value < t.value{
		t.left = add(t.left, value)
	} else {
		t.right = add(t.right, value)
	}

	return t
}