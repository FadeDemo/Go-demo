// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 73.

// Comma prints its argument numbers with a comma at each power of 1000.
//
// Example:
//
//	$ go build gopl.io/ch3/comma
//	$ ./comma 1 12 123 1234 1234567890
//	1
//	12
//	123
//	1,234
//	1,234,567,890
package main

import (
	"fmt"
	"os"
)

// test with 123456 7890.123 -987654.321 +123.456
func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("  %s\n", comma(os.Args[i]))
	}
}

// comma inserts commas in a non-negative decimal integer string.
// Enhanced to handle floating point numbers and optional signs.
func comma(s string) string {
	n := len(s)

	// 处理正负号
	sign := ""
	if s[0] == '+' || s[0] == '-' {
		sign = s[:1] // 保存符号
		s = s[1:]    // 去掉符号进行处理
		n = len(s)   // 更新长度
	}

	// 处理浮点数
	dotIndex := -1
	for i, char := range s {
		if char == '.' {
			dotIndex = i
			break
		}
	}

	if dotIndex != -1 {
		// 如果包含小数点，分别处理整数部分和小数部分
		intPart := comma(s[:dotIndex])   // 对整数部分加逗号
		fracPart := s[dotIndex:]         // 小数部分保持不变
		return sign + intPart + fracPart // 重新拼接
	}

	// 处理整数部分
	if n <= 3 {
		return sign + s
	}
	return sign + comma(s[:n-3]) + "," + s[n-3:]
}
