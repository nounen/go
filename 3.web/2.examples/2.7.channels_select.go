package main

import "fmt"

// TODO: 表示没看懂 select

// 如果存在多个channel的时候，我们该如何操作呢
// Go里面提供了一个关键字select通过select可以监听channel上的数据流动
func fibonacci(c, quit chan int) {
	x, y := 1, 1

	for {
		// select默认是阻塞的，只有当监听的channel中有发送或接收可以进行时才会运行，
		// 当多个channel都准备好的时候，select是随机的选择一个执行的 (TODO: 看着不像随机啊!!)
		select {
			case c <- x:
				x, y = y, x + y
			case <-quit:
				fmt.Println("quit")
				return
			// default: // 当c阻塞的时候执行这里
			// 	fmt.Println("default")
			// 	return
		}
	}
}

func main() {
	c    := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}

		quit <- 0
	}()

	fibonacci(c, quit)
}