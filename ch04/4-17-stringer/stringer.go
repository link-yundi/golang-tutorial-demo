package main

import "fmt"

/**
fmt包中定义的stringer是最普遍的接口之一
stringer是一个可以用字符串描述自己的类型，fmt包（还有很多包）都通过此接口来打印值
 */

type Person struct {
	Name string
	Age int
}

func (p Person) String() string{
	return fmt.Sprintf("%v (%v years)\n", p.Name, p.Age)
}

func main() {
	a := Person{"Arthur Dent", 42}
	z := Person{"Zaphod Beeblebrox", 9001}
	fmt.Println(a, z)
}
