package main

import "fmt"

/**
映射将键映射到值
映射的零值为nil。nil映射既没有键，也不能添加键
make函数会返回给定类型的映射，并将其初始化备用
 */

type Vertex struct {
	Lat, Long float64
}

var m map[string]Vertex

func main() {
	m = make(map[string]Vertex)
	m["Bell Labs"] = Vertex{
		40.32323,
		-74.5435,
	}
	fmt.Println(m["Bell Labs"])
}