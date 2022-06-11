package main

import "fmt"

var arr0 [5]int = [5]int{1, 2, 3}
var arr1 = [5]int{1, 2, 3, 4, 5}
var arr2 = [...]int{1, 2, 3, 4, 5, 6}
var str = [5]string{3: "hello", 4: "world"}

func modifyArr(a [2]int) {
	fmt.Printf("a:%p\n", &a)
	a[1] = 1000
}

func sumTest(a [5]int8, sum int8) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < i; j++ {
			if a[i]+a[j] == sum {
				fmt.Printf("(%d,%d)\n", i, j)
			}
		}
	}
}

func main() {
	a := [3]int{1, 2, 3}
	b := [...]int{1, 2, 3, 4}
	c := [5]int{2: 100, 4: 200}
	d := [...]struct {
		name string
		age  uint8
	}{
		{"user1", 10},
		{"user2", 20},
	}
	fmt.Println(arr0, arr1, arr2, str)
	fmt.Println(a, b, c, d)

	dd := [2]int{}
	fmt.Printf("a:%p\n", &dd)
	modifyArr(dd)
	fmt.Println(dd)

	ff := [2]int{}
	fmt.Println(len(ff), cap(ff))

	//    找出数组中和为给定值的两个元素的下标，例如数组[1,3,5,8,7]，
	// 找出两个元素之和等于8的下标分别是（0，4）和（1，2）

	sumArr := [5]int8{1, 3, 5, 8, 7}
	sumTest(sumArr, 8)
}
