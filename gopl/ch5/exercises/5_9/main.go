package main

import (
	"bytes"
	"fmt"
	"strings"
)

func main() {
	var data = "abc $foo ddd eee$foo"
	result := expand(data, test1)
	fmt.Println(result)
}

func expand(s string, f func(string) string) string {
	var buffer bytes.Buffer
	const sub = "$foo"
	i := 0
	for j := strings.Index(s[i:], sub); j != -1; j = strings.Index(s[i:], sub) {
		buffer.WriteString(s[i : i+j])
		buffer.WriteString(f("foo"))
		i += j + len(sub)
	}
	return buffer.String()
}

func test1(str string) string {
	return strings.ToUpper(str)
}
