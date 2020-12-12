package main

import (
	"fmt"
	"runtime"
)

// switch 是编写一连串 if-else 语句的简便方法。它运行的第一个值等于条件表达式的 case 语句
// go 的 switch 语句只运行选定的case，而非之后的所有case。实际上，go自动提供了在这些语言中每个case后面所需的break语句。
// 除非以fallthrough语句结束，否则分支会自动终止
// go的另一点重要的不同在于switch的case无需为常量，并且取值不必为整数


func main() {
	fmt.Print("Go runs on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OS X.")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s.\n", os)
	}
}
