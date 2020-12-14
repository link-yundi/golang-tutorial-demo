package main

import "fmt"

// 映射的文法与结构体相似，不过必须有键名

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex {
	"Bell Labs": Vertex {
		42.232,
		-89.4322,
	},
	"Google": Vertex {
		37.342,
		-122.342343,
	},
}

func main() {
	fmt.Println(m)
}