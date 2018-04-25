package main

import (
	"fmt"
	"math"
)

func test(x, y, lim float64) float64 {
	if v := math.Pow(x, y); v < lim {
		return v
	} else {
		// can use v here
		fmt.Printf("%g >= %g\n", v, lim)
	}
	// can't use v here
	return lim
}

func main() {
	fmt.Println(test(3, 2, 10))
	fmt.Println(test(3, 3, 2))
}
