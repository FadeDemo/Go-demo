package main

import "fmt"

func main() {
	var test = []int{1, 2, 3}
	for i := range test {
		fmt.Println(test[i])
	}
	for i, v := range test {
		fmt.Println(i, v)
	}
}
