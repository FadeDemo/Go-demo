// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 195.

// Http4 is an e-commerce server that registers the /list and /price
// endpoint by calling http.HandleFunc.
package main

import (
	"bytes"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
)

//!+main

func main() {
	db := database{"shoes": 50, "socks": 5}
	http.HandleFunc("/list", db.list)
	http.HandleFunc("/price", db.price)
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

//!-main

type dollars float32

func (d dollars) String() string { return fmt.Sprintf("$%.2f", d) }

type database map[string]dollars

func (db database) list(w http.ResponseWriter, req *http.Request) {
	var buffer bytes.Buffer
	buffer.WriteString("<table border=\"1\">\n")
	buffer.WriteString(`  <thead>
    <tr>
      <th>item</th>
      <th>price</th>
    </tr>
  </thead>
  <tbody>`)
	for item, price := range db {
		buffer.WriteString(`    <tr>
      <td>` + item + `</td>
      <td>` + fmt.Sprint(price) + `</td>
    </tr>`)
	}
	buffer.WriteString(`  </tbody>
</table>`)
	err := template.Must(template.New("list").Parse(buffer.String())).Execute(w, nil)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
	}
}

func (db database) price(w http.ResponseWriter, req *http.Request) {
	item := req.URL.Query().Get("item")
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
}
