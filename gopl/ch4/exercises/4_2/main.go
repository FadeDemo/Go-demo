// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 83.

// The sha256 command computes the SHA256 hash (an array) of a string.
package main

import (
	"bufio"
	"bytes"
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"os"
)

//!+

func main() {
	var sha = flag.String("sha", "256", "If you want sha256 hash")
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	var buffer bytes.Buffer
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break // 结束输入循环
		}
		buffer.WriteString(line)
	}
	if *sha == "256" {
		fmt.Printf("%x\n%[1]T\n", sha256.Sum256(buffer.Bytes()))
	} else if *sha == "384" {
		fmt.Printf("%x\n%[1]T\n", sha512.Sum384(buffer.Bytes()))
	} else if *sha == "512" {
		fmt.Printf("%x\n%[1]T\n", sha512.Sum512(buffer.Bytes()))
	} else {
		os.Exit(1)
	}
}

//!-
