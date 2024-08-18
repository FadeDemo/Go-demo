package main

import "fmt"

func unique(strings []string) []string {
	i := 0
	for j := 0; j < len(strings); j++ {
		if j == 0 || strings[j] != strings[j-1] {
			strings[i] = strings[j]
			i++
		}
	}
	return strings[:i]
}

func main() {
	//!+main
	data := []string{"one", "three", "three", "four"}
	fmt.Printf("%q\n", unique(data)) // `["one" "three"]`
	fmt.Printf("%q\n", data)         // `["one" "three" "three"]`
	//!-main
}

//!-alt
