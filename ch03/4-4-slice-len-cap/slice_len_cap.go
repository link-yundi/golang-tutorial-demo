package main

import "fmt"

/**
切片拥有长度和容量
切片的长度就是它所包含的元素个数
切片的容量是从**它**的第一个元素开始数，到其**底层数组**元素末尾的个数
切片 s 的长度和容量可以通过表达式len(s)和cap(s)来获取
你可以通过重新切片来扩展一个切片
 */

func main() {
	s := []int{2, 3, 5, 7, 11, 13}
	printSlice(s)

	// 截取切片使其长度为0
	s = s[:0]
	printSlice(s)

	// 拓展其长度
	s = s[:4]
	printSlice(s)

	// 舍弃前两个值
	s = s[2:]
	printSlice(s)
}

func printSlice(s []int) {
	fmt.Printf("len=%d cap=%d %v\n", len(s), cap(s), s)
}