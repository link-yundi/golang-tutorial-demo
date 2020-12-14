package main

/**
练习：映射
实现WordCount，它应当返回一个映射，其中包含字符串s中每个单词的个数。函数wc.Test会对此函数执行一系列测试用例，并输出成功还是失败
 */

import (
	"golang.org/x/tour/wc"
	"strings"
)

func WordCount(s string) map[string]int {
	splitString := strings.Split(s, " ")
	dict := make(map[string]int)
	for _, v := range splitString {
		if _, ok := dict[v]; ok {
			// 存在
			dict[v] += 1
		} else {
			dict[v] = 1
		}
	}
	return dict
}

func testSplit(s string) []string {
	return strings.Split(s, " ")
}

func main() {
	wc.Test(WordCount)
	//a := testSplit("I am learning Go!")
	//fmt.Println(a)

}
