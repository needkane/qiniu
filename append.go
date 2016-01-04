package main

import (
	"fmt"
	"os"
)

func main() {
	s1 := []string{"0", "1"}
	s1 = append(s1, "3")
	s1 = append(s1, "4")
	fmt.Println(s1)

}
