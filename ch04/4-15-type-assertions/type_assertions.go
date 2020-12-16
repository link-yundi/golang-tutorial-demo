package main

import (
	"fmt"
)

/**
类型断言
类型断言 提供了访问接口值底层具体值的方式
`t := i.(T)`
该语句断言接口值 i 保存了具体类型为 T，并将其底层类型为T的值赋予变量t
若 i 并未保存T类型的值，该语句就会触发一个panic
为了判断一个接口值是否保存了一个特定的类型，类型断言可以返回两个值：其底层值以及一个报告断言是否成功的布尔值
`t, ok := i.(T)`
若 i 保存了一个T，那么t将会是其底层值，而ok为true
否则，ok将为false而t将为T类型的零值，程序并不会产生panic
注意：这种语法和读取一个映射有相同之处
 */

func main() {
	var i interface{} = "hello"

	s := i.(string)
	fmt.Println(s)

	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	f = i.(float64)		// panic
	fmt.Println(f)
}
