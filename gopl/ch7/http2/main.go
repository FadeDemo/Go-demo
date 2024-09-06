// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 192.

// Http2 is an e-commerce server with /list and /price endpoints.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	log.Fatal(http.ListenAndServe("localhost:8000", db))
}

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

// !+handler
func (db database) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/list":
		for item, price := range db {
			_, err := fmt.Fprintf(w, "%s: %s\n", item, price)
			if err != nil {
				os.Exit(1)
			}
		}
	case "/price":
		item := req.URL.Query().Get("item")
		price, ok := db[item]
		if !ok {
			w.WriteHeader(http.StatusNotFound) // 404
			_, err := fmt.Fprintf(w, "no such item: %q\n", item)
			if err != nil {
				os.Exit(1)
			}
			return
		}
		_, err := fmt.Fprintf(w, "%s\n", price)
		if err != nil {
			os.Exit(1)
		}
	default:
		w.WriteHeader(http.StatusNotFound) // 404
		_, err := fmt.Fprintf(w, "no such page: %s\n", req.URL)
		if err != nil {
			os.Exit(1)
		}
	}
}

//!-handler
