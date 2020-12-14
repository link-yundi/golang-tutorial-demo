package main

import "fmt"

/**
练习：斐波那契闭包
	让我们用函数做一些好玩的事情
	实现一个fibonacci函数，它返回一个函数（闭包），该闭包返回一个斐波那契数列(0, 1, 1, 2, 3, 5)
 */

// 返回一个 闭包
func fibonacci() func() int {
	 a, b := 0, 1
	 indx := 1
	 return func() int {
	 	if indx == 1 {
	 		indx += 1
	 		return a
		} else {
			indx += 1
			a, b = b, a+b
			return a
		}
	 }
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}

