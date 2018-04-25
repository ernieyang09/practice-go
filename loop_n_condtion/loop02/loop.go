package main

import (
	"fmt"
	"math"
)

// Sqrt test function
func Sqrt(x float64) float64 {
	z := float64(500000)
	for i := 0; i < 2013; i++ {
		z = z - (z*z-x)/(2*z)
	}
	return z
}

func main() {
	i := float64(169)
	fmt.Println(Sqrt(i), Sqrt(i) == math.Sqrt(i))
}
