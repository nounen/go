package main

import (
	"fmt"
)

func main() {
	s := "\u00ff\u754c"

	for i, c := range s {
		fmt.Printf("%d:%c", i, c)
	}

	// 0:ÿ 2:界
}