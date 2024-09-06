// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 173.

// Bytecounter demonstrates an implementation of io.Writer that counts bytes.
package main

import (
	"fmt"
	"os"
)

//!+bytecounter

type ByteCounter int

func (c *ByteCounter) Write(p []byte) (int, error) {
	*c += ByteCounter(len(p)) // convert int to ByteCounter
	return len(p), nil
}

//!-bytecounter

func main() {
	//!+main
	var c ByteCounter
	_, err := c.Write([]byte("hello"))
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(c) // "5", = len("hello")

	c = 0 // reset the counter
	var name = "Dolly"
	_, err = fmt.Fprintf(&c, "hello, %s", name)
	if err != nil {
		os.Exit(1)
	}
	fmt.Println(c) // "12", = len("hello, Dolly")
	//!-main
}
