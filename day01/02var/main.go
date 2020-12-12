package main

import "fmt"

// 声明变量
//var name string
//var age int
//var idOk bool

// 批量声明
var (
	name string		// ""
	age int			// 0
	isOk bool		// false
)

func main() {
	name = "xxii"
	age = 16
	isOk = true
	// go语言中变量声明必须使用，不使用就编译不通过
	fmt.Printf("name:%s", name)	// %s: 占位符，使用name这个变量的值去替换占位符
	fmt.Print(isOk)						// 在终端输出要打印的内容，不换行
	fmt.Println(age)					// 打印完指定的内容之后会在后面加一个换行符

	// 声明变量同时赋值
	var s1 string = "zyd"
	fmt.Println(s1)
	// 类型推导（根据值判断该变量是什么类型）
	var s2 = "20"
	fmt.Println(s2)
	// 简短变量声明
	s3 := "hhhh"
	fmt.Println(s3)
}