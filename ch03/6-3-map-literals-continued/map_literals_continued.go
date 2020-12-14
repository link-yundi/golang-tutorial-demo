package main

import "fmt"

// 若顶级类型只是一个类型名，可以在文法的元素中省略它

type Vertex struct {
	Lat, Long float64
}

var m = map[string]Vertex{
	"Bell Labs": {40.4322, -74.3232},
	"Google": {37.3234, -122.3402},
}

func main() {
	fmt.Println(m)
}