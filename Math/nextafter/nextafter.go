package main

import (
	"fmt"
	"math"
)

func main() {
	// binary floating point
	fmt.Printf("Now you have %g problems.\n",
		math.Nextafter(2, 2))
	fmt.Printf("Now you have %g problems.\n",
		math.Nextafter(2, 3))
	fmt.Printf("Now you have %g problems.\n",
		math.Nextafter(2, 1))
}
