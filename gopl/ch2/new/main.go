package main

import (
	"fmt"
	"unsafe"
)

func main() {
	p := new(struct{})
	q := new([0]int)
	fmt.Println(uintptr(unsafe.Pointer(p)) == uintptr(unsafe.Pointer(q)))
}
