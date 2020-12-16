package main

import (
	"fmt"
	"math"
)

/**
方法：
	go没有类，不过可以为结构体类型定义方法
	方法就是一类带有特殊的 接收者 参数的函数
	方法接收者在它自己的参数列表内，位于 func 关键字和方法名之间
 */

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y * v.Y)
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs())
}

