package main

import (
	"fmt"
)

/**
底层值为nil的接口值
即便接口内的具体值为nil，方法仍然会被nil接收者调用
注意：保存了nil具体指的接口其自身并不是nil
 */

type I interface {
	M()
}

type T struct {
	S string
}

func (t *T) M()  {
	if t == nil {
		fmt.Println("<nil>")
		return
	}
	fmt.Println(t.S)
}

func main() {
	var i I
	var t *T
	i = t
	describe(i)
	i.M()

	i = &T{"Hello"}
	describe(i)
	i.M()
}

func describe(i I)  {
	fmt.Printf("(%v, %T)\n", i, i)
}