package main

import (
	"fmt"
	"math/rand"
)

func main() {

	for i := 0; i < 10; i++ {
		x := rand.Intn(4)
		fmt.Println(x)
	}
}
