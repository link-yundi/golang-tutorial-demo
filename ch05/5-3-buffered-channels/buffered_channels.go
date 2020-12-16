package main

import "fmt"

/*
带缓冲的信道：
信道可以是带有缓冲的，将缓冲长度作为第二个参数提供给make来初始化一个带缓冲的信道：
`ch := make(chan int, 100)`
注意：仅当信道的缓冲区填满之后，向其发送数据时才会阻塞。当缓冲区为空时，接收方才会阻塞。
*/

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}


