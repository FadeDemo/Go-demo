// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 10.
//!+

// Dup2 prints the count and text of lines that appear more than once
// in the input.  It reads from stdin or from a list of named files.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	//counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f)
			f.Close()
		}
	}
	//for line, n := range counts {
	//	if n > 1 {
	//		fmt.Printf("%d\t%s\n", n, line)
	//	}
	//}
}

func countLines(f *os.File) {
	counts := make(map[string]int)
	input := bufio.NewScanner(f)
	for input.Scan() {
		line := input.Text()
		if line == "" {
			break // 结束输入循环
		}
		counts[line]++
		if counts[line] > 1 {
			fmt.Println(f.Name())
			break
		}
	}
	// NOTE: ignoring potential errors from input.Err()
}

//!-
