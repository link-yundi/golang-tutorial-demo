package main

import "fmt"

/**
可以将下标或者值设置为 _ 来忽略它
 */

func main() {
	pow := make([]int, 10)
	for i := range pow {
		pow[i] = 1 << uint(i)
	}
	for _, value := range pow {
		fmt.Printf("%d \n", value)
	}
}
