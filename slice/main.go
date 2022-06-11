package main

import "fmt"

var arr = [...]int{0, 1, 2, 3, 4, 5, 6, 7, 8}
var (
	slice0 = arr[2:8]
	slice1 = arr[:6]
	slice2 = arr[5:]
	slice3 = arr[0:len(arr)]
	slice4 = arr[:len(arr)]
)

var slice00 []int = make([]int, 10)
var slice01 = make([]int, 10)
var slice02 = make([]int, 10, 10)

func main() {
	/*var s1 []int
	if s1 == nil {
		fmt.Println("s1为空")
	} else {
		fmt.Println("s1不为空")
	}

	s2 := []int{}
	var s3 []int = make([]int, 0)
	fmt.Println(s1, s2, s3)

	var s4 []int = make([]int, 0, 0)
	fmt.Println(s4)

	s5 := []int{1, 2, 3, 4, 5}
	fmt.Println(s5)

	arr6 := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	s6 = arr6[1:4]
	fmt.Println(s6)
	s7 := arr6[:len(arr6)]
	fmt.Println(s7)*/

	slice000 := make([]int, 10)
	slice001 := make([]int, 10)
	slice002 := make([]int, 10, 10)

	fmt.Printf("全局变量arr：%v\n", arr)
	fmt.Printf("全局变量slice0：%v\n", slice0)
	fmt.Printf("全局变量slice1：%v\n", slice1)
	fmt.Printf("全局变量slice2：%v\n", slice2)
	fmt.Printf("全局变量slice3：%v\n", slice3)
	fmt.Printf("全局变量slice4：%v\n", slice4)

	fmt.Printf("make全局变量slice00：%v\n", slice00)
	fmt.Printf("make全局变量slice01：%v\n", slice01)
	fmt.Printf("make全局变量slice02：%v\n", slice02)
}
