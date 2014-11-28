package main

import (
	"fmt"
	"os"
)

func main() {
	_, err := os.Open("ff.jpg")
	fmt.Println(err.Error())
}
