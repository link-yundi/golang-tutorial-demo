package main

import "fmt"

/**
nil 切片
切片的零值是nil
nil切片的长度和容量为0并且没有底层数组
 */

func main() {
	var s []int
	fmt.Println(s, len(s), cap(s))
	if s == nil {
		fmt.Println("nil!")
	}
}
