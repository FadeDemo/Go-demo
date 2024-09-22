package main

import (
	"fmt"
	"os"
	"strings"
	"time"
)

func main() {
	fmt.Printf("EchoWithLowPerformance: %.10fs elapsed\n", EchoWithLowPerformance())
	fmt.Printf("EchoUsingStringsJoin: %.10fs elapsed\n", EchoUsingStringsJoin())
}

func EchoWithLowPerformance() float64 {
	start := time.Now()
	s, sep := "", ""
	for _, arg := range os.Args[1:] {
		s += sep + arg
		sep = " "
	}
	return time.Since(start).Seconds()
}

func EchoUsingStringsJoin() float64 {
	start := time.Now()
	strings.Join(os.Args[1:], " ")
	return time.Since(start).Seconds()
}
