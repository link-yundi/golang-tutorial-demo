package main

import (
	"fmt"
	"sync"
	"time"
)

/**
channel 非常适合在各个go程之间进行通信
如果我们并不需要通信。比如说，我们只是想保证每次只有一个go程能够访问一个共享的变量，从而避免冲突
这里涉及的概念叫做"互斥"，我们通常使用互斥锁"Mutex"这一个数据结构来提供这种机制
go 标准库中提供了`sync.Mutex`互斥锁类型以及两个方法
`Lock Unlock`
我们可以通过在代码前调用Lock方法，在代码后面调用Unlock方法来保证一段代码的互斥执行
我们也可以用defer语句来保证互斥锁一定会被解锁
 */

// SafeCounter的并发使用是安全的
type SafeCounter struct {
	v map[string]int
	mux sync.Mutex
}

// Inc 增加给定key的计数器的值
func (c *SafeCounter) Inc(key string) {
	c.mux.Lock()
	// Lock之后同一时刻只有一个goroutine能够访问c.v
	defer c.mux.Unlock()
	c.v[key]++
}

// Value 返回给定 key 的计数器的当前值
func (c *SafeCounter) Value(key string) int {
	c.mux.Lock()
	defer c.mux.Unlock()
	return c.v[key]
}

func main() {
	c := SafeCounter{v: make(map[string]int)}
	for i := 0; i < 1000; i++ {
		go c.Inc("somekey")
	}
	time.Sleep(time.Second)
	fmt.Println(c.Value("somekey"))
}