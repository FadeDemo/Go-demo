// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"sync"
)

var mu sync.Mutex

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	http.HandleFunc("/update", db.update)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	mu.Lock()
	for item, price := range db {
		_, err := fmt.Fprintf(w, "%s: %s\n", item, price)
		if err != nil {
			os.Exit(1)
		}
	}
	defer mu.Unlock()
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	if price, ok := db[item]; ok {
		_, err := fmt.Fprintf(w, "%s\n", price)
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
	defer mu.Unlock()
}

func (db database) update(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
	mu.Lock()
	if _, ok := db[item]; !ok {
		w.WriteHeader(http.StatusBadRequest)
		_, _ = fmt.Fprintf(w, "no such item: %q\n", item)
		return
	} else {
		price := req.URL.Query().Get("price")
		if price == "" {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "price cannot be null\n")
			return
		}
		value, err := strconv.ParseFloat(price, 32)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			_, _ = fmt.Fprintf(w, "wrong price format: %s\n", price)
			return
		}
		db[item] = dollars(value)
	}
	defer mu.Unlock()
}
