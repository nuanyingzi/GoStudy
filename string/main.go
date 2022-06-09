package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	/*str1 := "I want to study Golang"
	count := len(str1)
	fmt.Println(count)

	str2 := "And you?"

	str3 := fmt.Sprintf("%s. %s", str1, str2)
	fmt.Println(str3)

	// 分割字符串
	str3List := strings.Split(str3, ".")
	fmt.Printf("str3List类型:%T 值:%v\n", str3List, str3List)

	// 判断是否包括
	boolStr3 := strings.Contains(str3, "want")
	fmt.Printf("boolStr3类型:%T 值:%v\n", boolStr3, boolStr3)

	// 判断前缀
	str3Prefix := strings.HasPrefix(str3, "want")
	fmt.Printf("str3Prefix类型:%T 值:%v\n", str3Prefix, str3Prefix)

	// 判断后缀
	str3Suffix := strings.HasSuffix(str3, "you?")
	fmt.Printf("str3Suffix类型:%T 值:%v\n", str3Suffix, str3Suffix)

	// 子串出现的位置
	str3Index := strings.Index(str3, "want")
	fmt.Printf("str3Index类型:%T 值:%v\n", str3Index, str3Index)

	// 子串最后出现的位置
	str3LastIndex := strings.LastIndex(str3, "o")
	fmt.Printf("str3LastIndex类型:%T 值:%v\n", str3LastIndex, str3LastIndex)*/

	num := 50000
	list := make([]string, 100)
	for i := 0; i < num; i++ {
		list = append(list, "Hello world")
	}
	var test1, test2 string
	start1 := time.Now()
	for i := 0; i < num; i++ {
		test1 += list[i] + " "
	}
	end1 := time.Now().Sub(start1)
	start2 := time.Now()
	test2 = strings.Join(list, " ")
	end2 := time.Now().Sub(start2)
	fmt.Printf("%v\n %v\n", test1, test2)
	fmt.Println(end1, end2)
	// 差距有点大哦
	// 11.8929295s 2.9981ms
}
