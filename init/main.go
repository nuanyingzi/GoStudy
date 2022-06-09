package main

import "fmt"

// 同一个Go文件inir函数按顺序执行

func init() {
	fmt.Println("This is init01")
}

func init() {
	fmt.Println("This is init02")
}

func init() {
	fmt.Println("This is init03")
}

func main() {
	fmt.Println("初始化完成")
}
