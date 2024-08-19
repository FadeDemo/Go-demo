package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countWords(os.Stdin, counts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				_, _ = fmt.Fprintf(os.Stderr, "wordfreq: %v\n", err)
				continue
			}
			countWords(f, counts)
			err = f.Close()
			if err != nil {
				continue
			}
		}
	}
	for word, n := range counts {
		fmt.Printf("[%s]=%d\n", word, n)
	}
}

func countWords(f *os.File, counts map[string]int) {
	input := bufio.NewScanner(f)
	input.Split(bufio.ScanWords)
	for input.Scan() {
		counts[input.Text()]++
	}
}
