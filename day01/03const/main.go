package main

import "fmt"

const pi = 3.1415926
// 批量声明常量
const (
	statusOK = 200
	notFound = 404
)
// 批量声明常量时，如果某一行声明后没有赋值，默认和上一行一致
const (
	n1 = 100
	n2
	n3
)

func main() {
	fmt.Println("n1: ", n1)
	fmt.Println("n2: ", n2)
	fmt.Println("n3: ", n3)
}
