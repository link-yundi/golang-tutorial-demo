package main

import (
	"fmt"
	"time"
)

/**
go 程(goroutines) 是go运行时管理的轻量级线程
`go f(x, y, z)`
会启动一个新的go程并执行
f(x, y, x)
f, x, y和z的求值发生在当前的go程中，而f的执行发生在新的go程中
go程在相同的地址空间中运行，因此在访问共享的内存时必须进行同步。sync包提供了这种能力，不过在go中并不经常用到，因为还有其他的方法
 */

func say(s string) {
	for i := 0; i < 5; i++ {
		time.Sleep(100 * time.Millisecond)
		fmt.Println(s)
	}
}

func main() {
	go say("hello")
	say("world")
}


