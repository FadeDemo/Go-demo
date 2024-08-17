package main

import (
	"fmt"
)

func main() {
	s := "你好"
	fmt.Println(s)
	fmt.Println(len(s))
	fmt.Printf("%c\n", s[0])

	a := 'a'
	b := '你'
	fmt.Printf("%c\n", a)
	fmt.Printf("%c\n", b)
	fmt.Printf("% x\n", s)
	r := []rune(s)
	fmt.Printf("%x\n", r)
	r = []int32(s)
	fmt.Printf("%x\n", r)
	for i, v := range "abc" {
		fmt.Printf("%d %T\n", i, v)
	}
}
