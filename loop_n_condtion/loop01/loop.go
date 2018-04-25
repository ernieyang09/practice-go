package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10000000; i++ {
		sum += i
	}

	fmt.Println(sum)

	// for continue while
	sum2 := 1

	for sum2 < 1000 {
		sum2 += sum2
	}

	fmt.Println(sum2)

	// infinity
	// for {
	// }
}
