package main

// import 用_表示引用这个包只为了调用包内的init函数，无法调用包内其他函数

import (
	_ "GoStudy/import/hello"
	"fmt"
)

func main() {
	//hello.Hello()
	fmt.Println("This is import main()")
}
