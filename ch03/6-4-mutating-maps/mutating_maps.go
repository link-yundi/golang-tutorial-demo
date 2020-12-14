package main

import "fmt"

/**
修改映射
在映射m中插入或者修改元素：`m[key] = elem`
获取元素：`elem = m[key]`
删除元素：`delete(m, key)`
通过双赋值检测某个键是否存在：elem, ok = m[key]
如果key在m中，ok为true，否则，ok为false
如果key不在映射中，那么elem是该映射元素类型的零值
同样的，当从映射中读取某个不存在键时，结果是映射的元素类型的零值
注意：若elem或ok还未声明，可以使用短变量声明：elem, ok := m[key]
 */

func main() {
	m := make(map[string]int)

	m["Answer"] = 42
	fmt.Println("The value:", m["Answer"])

	m["Answer"] = 48
	fmt.Println("The value:", m["Answer"])

	delete(m, "Answer")
	fmt.Println("The value:", m["Answer"])

	v, ok := m["Answer"]
	fmt.Println("The value:", v, "Present?", ok)
}