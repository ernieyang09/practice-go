package main

import "fmt"

var i int
var test, test2 bool
var j, k int = 2, 3
var test3, test4 = "ss", true

// haha := "hahaha"

func main() {
	// only in function block
	haha := "hahaha"
	fmt.Println(i, test, test2)
	fmt.Println(j, k)
	fmt.Println(test3, test4)
	fmt.Println(haha)
}
