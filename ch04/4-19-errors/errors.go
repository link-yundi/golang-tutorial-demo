package main

import (
	"fmt"
	"time"
)

/**
go 程序使用error值来表示错误状态
与fmt.Stringer类似，error类型是一个内建接口
`type error interface {
	Error()	string
}`
通常函数会返回一个error值，调用它的代码应当判断这个错误是否等于nil来进行错误处理
 */

type MyError struct {
	When time.Time
	What string
}

// 实现了Error()方法，是error接口的一个实现
func (e *MyError) Error() string {
	return fmt.Sprintf("at %v, %s", e.When, e.What)
}

func run() error {
	return &MyError{
		time.Now(),
		"it didn't work",
	}
}

func main() {
	if err := run(); err != nil {
		fmt.Println(err)
	}
}