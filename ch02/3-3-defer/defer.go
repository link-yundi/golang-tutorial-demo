package main

// defer语句会将函数推迟到外层函数返回之后执行
// 推迟调用的函数其参数会立即求值，但是直到外层函数返回前该函数都不会被调用

import "fmt"

func main() {
	defer fmt.Println("world")
	fmt.Println("hello")
}
