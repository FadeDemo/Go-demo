// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// Run with "web" command-line argument for web server.
// See page 13.
//!+main

// Lissajous generates GIF animations of random Lissajous figures.
package main

import (
	"log"
	"math/rand"
	"net/http"
	"os"
	"time"
)

//!-main
// Packages not needed by version in book.

//!+main

func main() {
	// Note: Do not run directly, redirect the output to a file in CMD or Powershell 7.x
	//!-main
	// The sequence of images is deterministic unless we seed
	// the pseudo-random number generator using the current time.
	// Thanks to Randall McPherson for pointing out the omission.
	//rand.Seed(time.Now().UTC().UnixNano())
	rand.New(rand.NewSource(time.Now().UTC().UnixNano()))

	if len(os.Args) > 1 && os.Args[1] == "web" {
		//!+http
		handler := func(w http.ResponseWriter, r *http.Request) {
			//lissajous(w)
			Generate(w)
		}
		http.HandleFunc("/", handler)
		//!-http
		log.Fatal(http.ListenAndServe("localhost:8000", nil))
		return
	}
	//fmt.Println("你好")
	//!+main
	//lissajous(os.Stdout)
	Generate(os.Stdout)
}

//!-main
