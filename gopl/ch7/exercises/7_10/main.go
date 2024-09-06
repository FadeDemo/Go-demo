package main

import (
	"fmt"
	"sort"
)

func main() {
	var data = "hello"
	fmt.Printf("Whether data %s is palindrome:%t\n", data, IsPalindrome(RuneSlice(data)))
	data = "aba"
	fmt.Printf("Whether data %s is palindrome:%t\n", data, IsPalindrome(RuneSlice(data)))
	data = "12345654321"
	fmt.Printf("Whether data %s is palindrome:%t\n", data, IsPalindrome(RuneSlice(data)))
	data = "987654321"
	fmt.Printf("Whether data %s is palindrome:%t\n", data, IsPalindrome(RuneSlice(data)))
	data = "madam"
	fmt.Printf("Whether data %s is palindrome:%t\n", data, IsPalindrome(RuneSlice(data)))
}

// RuneSlice 定义一个自定义类型 RuneSlice 来表示字符串的 rune 切片
type RuneSlice []rune

// Len 实现 sort.Interface 接口的 Len 方法
func (r RuneSlice) Len() int {
	return len(r)
}

// Less 实现 sort.Interface 接口的 Less 方法
func (r RuneSlice) Less(i, j int) bool {
	// 判断第 i 个字符是否小于第 j 个字符
	return r[i] < r[j]
}

// Swap 实现 sort.Interface 接口的 Swap 方法
func (r RuneSlice) Swap(i, j int) {
	r[i], r[j] = r[j], r[i]
}

func IsPalindrome(s sort.Interface) bool {
	for i, j := 0, s.Len()-1; i < j; i, j = i+1, j-1 {
		if s.Less(i, j) || s.Less(j, i) {
			return false
		}
	}
	return true
}
