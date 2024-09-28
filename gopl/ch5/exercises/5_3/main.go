// Copyright © 2016 Alan A. A. Donovan & Brian W. Kernighan.
// License: https://creativecommons.org/licenses/by-nc-sa/4.0/

// See page 122.
//!+main

// Findlinks1 prints the links in an HTML document read from standard input.
package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

// go run ch1/fetch/main.go https://golang.org | go run ch5/exercises/5_3/main.go
func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "findlinks1: %v\n", err)
		os.Exit(1)
	}
	printText(doc)
}

//!-main

// printText appends to links each link found in n and returns the result.
func printText(n *html.Node) {
	if n.Type == html.TextNode {
		// 输出文本节点的内容
		fmt.Println(n.Data)
	}

	// 如果遇到 <script> 或 <style> 元素，直接跳过它及其子节点的遍历
	if n.Type == html.ElementNode && (n.Data == "script" || n.Data == "style") {
		return
	}

	// 递归遍历子节点
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		printText(c)
	}
}

/*
//!+html
package html

type Node struct {
	Type                    NodeType
	Data                    string
	Attr                    []Attribute
	FirstChild, NextSibling *Node
}

type NodeType int32

const (
	ErrorNode NodeType = iota
	TextNode
	DocumentNode
	ElementNode
	CommentNode
	DoctypeNode
)

type Attribute struct {
	Key, Val string
}

func Parse(r io.Reader) (*Node, error)
//!-html
*/
