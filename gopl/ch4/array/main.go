package main

import "fmt"

type Currency int

const (
	USD Currency = iota
	EUR
	GBP
	RMB
)

func main() {
	symbols := [...]string{USD: "$", EUR: "€", GBP: "£", RMB: "¥"}
	fmt.Println(RMB, symbols[RMB], symbols[0], symbols[2])
	var arr [32]int
	test(arr)
	fmt.Println(arr[0])
}

func test(arr [32]int) {
	arr[0] = 9
}
