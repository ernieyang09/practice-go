package main

import "fmt"

func test() (result int, boolean bool) {
	return 1, false
}

func main() {
	// blank identifier
	_, error := test()
	fmt.Println(error)
}
