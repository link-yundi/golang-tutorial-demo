package main

import (
	"fmt"
	"math"
)

/**
接口类型是由一组方法签名定义的集合
接口类型的变量可以保存任何实现了这些方法的值
 */

type Abser interface {
	Abs() float64
}

type MyFloat float64

func (f MyFloat) Abs() float64 {
	if f < 0 {
		return float64(-f)
	}
	return float64(f)
}

type Vertex struct {
	X, Y float64
}

func (v *Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	var a Abser
	f := MyFloat(-math.Sqrt2)
	v := Vertex{3, 4}

	a = f	// a MyFloat 实现了 Abser
	a = &v 	// a *Vertex 实现了 Abser

	// a = v 	// 不是 *Vertex 没有实现 Abser
	fmt.Println(a.Abs())
}

// 思考：在go的世界里，万物都是接口，接口的实现方式就是实现其中的方法。MyFloat/*Vertex 由于都实现Abs()方法，所以它们都是Abser接口的一种实现

