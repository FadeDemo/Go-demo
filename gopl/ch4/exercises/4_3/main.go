// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 86.

// Rev reverses a slice.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//!+array
	a := [...]int{0, 1, 2, 3, 4, 5}
	//reverse(a[:])
	reverse(&a, 0, len(a)-1)
	fmt.Println(a) // "[5 4 3 2 1 0]"
	//!-array

	//!+slice
	s := [...]int{0, 1, 2, 3, 4, 5}
	// Rotate s left by two positions.
	//reverse(s[:2])
	reverse(&s, 0, 1)
	//reverse(s[2:])
	reverse(&s, 2, len(a)-1)
	reverse(&s, 0, len(a)-1)
	fmt.Println(s) // "[2 3 4 5 0 1]"
	//!-slice

	// Interactive test of reverse.
	input := bufio.NewScanner(os.Stdin)
outer:
	for input.Scan() {
		var ints [6]int
		for i, s := range strings.Fields(input.Text()) {
			x, err := strconv.ParseInt(s, 10, 64)
			if err != nil {
				_, err := fmt.Fprintln(os.Stderr, err)
				if err != nil {
					os.Exit(1)
				}
				continue outer
			}
			if i >= 6 {
				continue outer
			}
			ints[i] = int(x)
		}
		reverse(&ints, 0, len(ints)-1)
		fmt.Printf("%v\n", ints)
	}
	// NOTE: ignoring potential errors from input.Err()
}

// !+rev
// reverse reverses a slice of ints in place.
func reverse(s *[6]int, i, j int) {
	for ; i < j; i, j = i+1, j-1 {
		(*s)[i], (*s)[j] = (*s)[j], (*s)[i]
	}
}

//!-rev
