package main

import "fmt"

const (
	KB = 1000      // 1000 的 1 次方
	MB = KB * 1000 // 1000 的 2 次方
	GB = MB * 1000 // 1000 的 3 次方
	TB = GB * 1000 // 1000 的 4 次方
	PB = TB * 1000 // 1000 的 5 次方
	EB = PB * 1000 // 1000 的 6 次方
	//ZB = EB * 1000 // 1000 的 7 次方
	//YB = ZB * 1000 // 1000 的 8 次方
)

func main() {
	fmt.Println("1 KB =", KB, "bytes")
	fmt.Println("1 MB =", MB, "bytes")
	fmt.Println("1 GB =", GB, "bytes")
	fmt.Println("1 TB =", TB, "bytes")
	fmt.Println("1 PB =", PB, "bytes")
	fmt.Println("1 EB =", EB, "bytes")
	// will overflow
	//fmt.Println("1 ZB =", ZB, "bytes")
	//fmt.Println("1 YB =", YB, "bytes")
}
