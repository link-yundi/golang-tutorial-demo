package main

/**
等价二叉查找树
不同二叉树的叶节点上可以保存相同的值序列。
1. 实现walk函数
2. 测试walk函数
	函数`tree.New(k)`用于构造一个随机结构的已排序二叉查找树，它保存了值k,2k,3k,...,10k
	创建一个新的信道ch并且对其进行步进：
	`go Walk(tree.New(1), ch)`
	然后从信道中读取并且打印10个值，应当是数字1, 2, 3,..., 10
3. 用Walk实现Same函数来检测t1和t2是否存储了相同的值
4. 测试Same函数
	`Same(tree.New(1), tree.New(1))` 应当返回true，而Same(tree.New(1), tree.New(2))应当返回false
 */


