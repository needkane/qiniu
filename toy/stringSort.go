package main

import (
	"log"
	"os"
)

func main() {
	input := os.Args[1]
	log.Println(input)
	if len(input) < 2 || len(input) > 1000 {
		log.Println("sorry,2<= (len) <=  1000")
	}
	mid := len(input)/2 + 1
	//alphabet 65-90 97-122
	for i := 0; i < len(input)-2; i++ {
		for j := 1; j < len(input)-1; j++ {
			if !IfAlphabet(input[i]) {
				temp = input[i]
			}
		}
	}
}

func IfAlphabet(i byte) bool {
	if (i >= 65 && i <= 90) || (i >= 97 && i <= 122) {
		return true
	}
	return false
}
