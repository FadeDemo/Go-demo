// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 97.
//!+

// Charcount computes counts of Unicode characters.
package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"unicode"
)

func main() {
	//counts := make(map[rune]int)
	mapping := map[int]string{0: "letter(s)", 1: "number(s)", 2: "other Unicode character(s)"}
	var counts [3]int
	in := bufio.NewReader(os.Stdin)
	for {
		r, n, err := in.ReadRune() // returns rune, nbytes, error
		if err == io.EOF {
			break
		}
		if err != nil {
			_, _ = fmt.Fprintf(os.Stderr, "charcount: %v\n", err)
			os.Exit(1)
		}
		if r == unicode.ReplacementChar && n == 1 {
			continue
		}
		if unicode.IsLetter(r) {
			counts[0]++
		} else if unicode.IsNumber(r) {
			counts[1]++
		} else {
			counts[2]++
		}
	}
	for i, val := range counts {
		fmt.Printf("There is(are) %d %s.\n", val, mapping[i])
	}
}

//!-
