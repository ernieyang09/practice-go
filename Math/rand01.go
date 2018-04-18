package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Printf("This is fixed random %d\n", rand.Intn(100))
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("This is fixed random %d\n", rand.Intn(100))
	rand.Seed(time.Now().UnixNano())
	fmt.Printf("This is fixed random %d\n", rand.Intn(100))
}
