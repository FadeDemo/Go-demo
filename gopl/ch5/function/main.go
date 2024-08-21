package main

import "fmt"

func main() {

}

func add(x, y int) int {
	return x + y
}

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = fmt.Errorf("division by zero")
		return
	}
	result = a / b
	return
}
