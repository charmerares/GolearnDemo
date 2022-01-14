package main

import "fmt"

func main() {
	// 普通数组
	var a [3]int
	fmt.Println(a[0])

	for i, j := range a {
		fmt.Println(i, j)
	}

	i := 10
	fmt.Println(i)
	for i, j := range a {
		fmt.Println(i, j)
	}

	var q [3]int = [3]int{1, 2, 3}
	fmt.Println(q[1])

	p := [...]int{4, 5, 6}
	fmt.Println(p[2])

	// 多维数组
	// 声明一个二维整型数组，两个维度的长度分别是 4 和 2
	var array [4][2]int
	// 使用数组字面量来声明并初始化一个二维整型数组
	array = [4][2]int{{10, 11}, {20, 21}, {30, 31}, {40, 41}}
	// 声明并初始化数组中索引为 1 和 3 的元素
	array = [4][2]int{1: {20, 21}, 3: {40, 41}}
	// 声明并初始化数组中指定的元素
	array = [4][2]int{1: {0: 20}, 3: {1: 41}}

	fmt.Println(array[1][1])
}
