// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http3a is an e-commerce server that registers the /list and /price
// endpoints by calling (*http.ServeMux).HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func main() {
	db := database{"shoes": 50, "socks": 5}
	mux := http.NewServeMux()
	//!+main
	mux.HandleFunc("/list", db.list)
	mux.HandleFunc("/price", db.price)
	//!-main
	log.Fatal(http.ListenAndServe("localhost:8000", mux))
}

type database map[string]int

func (db database) list(w http.ResponseWriter, req *http.Request) {
	for item, price := range db {
		_, err := fmt.Fprintf(w, "%s: $%d\n", item, price)
		if err != nil {
			os.Exit(1)
		}
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	if price, ok := db[item]; ok {
		_, err := fmt.Fprintf(w, "$%d\n", price)
		if err != nil {
			os.Exit(1)
		}
	} else {
		w.WriteHeader(http.StatusNotFound) // 404
		_, err := fmt.Fprintf(w, "no such item: %q\n", item)
		if err != nil {
			os.Exit(1)
		}
	}
}

/*
//!+handlerfunc
package http

type HandlerFunc func(w ResponseWriter, r *Request)

func (f HandlerFunc) ServeHTTP(w ResponseWriter, r *Request) {
	f(w, r)
}
//!-handlerfunc
*/
