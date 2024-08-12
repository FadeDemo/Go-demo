// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 45.

// (Package doc comment intentionally malformed to demonstrate golint.)
// !+
package popcount

// pc[i] is the population count of i.
var pc [256]byte

func init() {
	for i := range pc {
		// pc[i/2]代表前n-1位有几个1，byte(i&1)判断最后一位是不是1
		pc[i] = pc[i/2] + byte(i&1)
	}
}

// PopCount returns the population count (number of set bits) of x.
// 这个函数存在的必要性是x是uint64类型的数，而pc只能表示一个字节范围的数
func PopCount(x uint64) int {
	var count byte = 0
	for i := 0; i < 8; i++ {
		count += pc[byte(x>>(i*8))]
	}
	return int(count)
}

//!-
