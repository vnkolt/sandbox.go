package main

import (
	"fmt"
)

func BitCount(x uint64) int {
	var res int
	for x != 0 { // almost like while( x != 0 ) in C/C++
		res++
		x &= x - 1
	}
	return res
}

func main() {
	var x uint64
	fmt.Printf("Input integer number x:")
	fmt.Scanf("%d", &x)
	fmt.Printf("x = %d(%b), BitCount(%d) = %d", x, x, x, BitCount(x)) // %b means base 2, e.g. 8 will be printed like 1000
}
