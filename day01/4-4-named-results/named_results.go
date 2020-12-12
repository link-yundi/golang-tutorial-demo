package main

import "fmt"

// 命名返回值
// go的返回值可以被命名，它们会被是做定义在函数顶部的变量
// 返回值的名称应当具有一定的意义，它可以作为文档使用
// 没有参数的return语句返回已命名的返回值。也就是直接返回

func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum -x
	return
}

func main() {
	fmt.Println(split(17))
}
