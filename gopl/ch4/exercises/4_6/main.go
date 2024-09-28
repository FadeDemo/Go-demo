package main

import (
	"fmt"
	"unicode"
	"unicode/utf8"
)

func squashSpaces(b []byte) []byte {
	// 用于记录处理后的 slice 长度
	w := 0
	// 标记是否处于空白字符的序列中
	inSpace := false

	for i := 0; i < len(b); {
		// 读取下一个 UTF-8 字符
		r, size := utf8.DecodeRune(b[i:])

		if unicode.IsSpace(r) {
			if !inSpace {
				// 如果这是第一个空白字符，写入一个空格
				b[w] = ' '
				w++
				inSpace = true
			}
			// 跳过这个空白字符
		} else {
			// 非空白字符，写入到当前的位置
			// 担心copy不是就地算法的话，可以用for循环代替
			copy(b[w:], b[i:i+size])
			w += size
			inSpace = false
		}

		// 移动到下一个字符
		i += size
	}

	return b[:w]
}

func main() {
	// 示例测试
	b := []byte("Hello,\t世界!\n  How \r\nare you?")
	fmt.Printf("Before: %s\n", string(b))

	// 调用函数
	b = squashSpaces(b)
	fmt.Printf("After:  %q\n", string(b))
}
