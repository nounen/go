package main

import (
	"os"
	"fmt"
)

// 命令行参数
// `go run 1.2.1.command.go canshu1 canshu2`
func main() {
	var s, sep string

	for i := 0; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}

	fmt.Println(s)
}
