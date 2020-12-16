package main

import "fmt"

/**
发送者可以通过 close 关闭一个信道来表示没有需要发送的值了。接收者可以通过为接收表达式分配
第二个参数来测试信道是否被关闭；若没有值可以接受且信道已经被关闭，那么在执行完
`v, ok := <-ch`
之后，ok会被设置为false
循环 `for i := range c` 会不断从信道接收值，直到它被关闭
注意：只有发送者才能关闭信道，接收者不能关闭。向一个已经关闭的信道发送数据会引发程序panic
还要注意：信道与文件不一样，通常情况下不用关闭它。只有在必须告诉接收者不再有需要发送的值时才有必要关闭，
例如终止一个range循环
 */

func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func main() {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for item := range c {
		fmt.Println(item)
	}
}
