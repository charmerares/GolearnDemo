package main

import (
	"fmt"
)

/* 注释打开会报错
c:=20
*/

var d string = "hello"

func main() {
	// 普通申请变量
	var a string = "hello world"
	fmt.Println(a)

	// 简短声明,此处是在mian函数中，
	b := 20
	fmt.Println(b)

	// 多变量同时赋值
	e, f := 10, 20
	fmt.Println(e, f)
	// 只需要保证只有一个变量未被初始化
	b, g := 10, 20
	fmt.Println(b, g)

	// 还可以"抛弃"部分返回值，即使用匿名变量
	_, h := getData()
	fmt.Println(h)

	const (
		FlagNone = 1 << iota
		FlagRed
		FlagGreen
		FlagBlue
	)
	fmt.Printf("%d %d %d\n", FlagRed, FlagGreen, FlagBlue)
}

func getData() (int, int) {
	return 10, 20
}
