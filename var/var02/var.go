package main

import "fmt"

// MAX hahaha
var (
	test       = true
	i          = 23
	MAX  int32 = 1<<31 - 1
	MIN        = ^MAX
)

func main() {
	const t = "%T(%v)\n"

	fmt.Printf(t, test, test)
	fmt.Printf(t, i, i)
	fmt.Printf(t, MAX, MAX)
	fmt.Printf(t, MIN, MIN)
}
