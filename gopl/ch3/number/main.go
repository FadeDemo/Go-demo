package main

import (
	"fmt"
	"math"
)

func main() {
	//i := 0.0 / 0
	//fmt.Println(math.IsNaN(i))
	//const j = 0.0
	//const k = 0
	//fmt.Println(math.IsNaN(j / k))
	i := 0.0 / 0.0
	fmt.Println(math.IsNaN(i))
}
