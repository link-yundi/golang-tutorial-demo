package main

import (
	"fmt"
	"time"
)

/**
默认选择：
当select中的其他分支都没有准备好时，default分支就会执行
为了在尝试发送或者接收时不发生阻塞，可以使用default分支
 */

func main() {
	tick := time.Tick(100 * time.Millisecond)
	boom := time.After(500 * time.Millisecond)
	for {
		select {
		case <- tick:
			fmt.Println("tick.")
		case <- boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("       .")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
