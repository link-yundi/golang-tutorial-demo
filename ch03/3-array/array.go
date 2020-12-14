package main

import "fmt"

/**
	类型 [n]T表示拥有n个T类型的值的数组
	表达式
	var a [10]int		将变量a声明为拥有10个整数的数组
	数组的长度是其类型的一部分，因此数组不能改变大小。
 */

func main() {
	var a [2]string
	a[0] = "Hello"
	a[1] = "World"
	fmt.Println(a[0], a[1])
	fmt.Println(a)

	primes := [6]int{1, 2, 3, 4}
	fmt.Println(primes)
	fmt.Println(primes[0], primes[5])

	b := []int{1, 3, 4}
	fmt.Println(b)

}
