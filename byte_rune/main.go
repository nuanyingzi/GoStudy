package main

import "fmt"

func main() {
	s := "prof.博客"
	for i := 0; i < len(s); i++ {
		fmt.Printf("%v(%c) ", s[i], s[i])
	}
	fmt.Println()

	for _, r := range s {
		fmt.Printf("%v(%c) ", r, r)
	}
	fmt.Println()

	s1 := "hello"
	byteS1 := []byte(s1)
	byteS1[0] = 'H'
	fmt.Println(string(byteS1))

	s2 := "博客"
	runeS2 := []rune(s2)
	runeS2[0] = '说'
	fmt.Println(string(runeS2))
}
