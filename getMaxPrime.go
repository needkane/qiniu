package main

import (
	"fmt"
)

func main() {

	fmt.Println("Enter a number")
	var i, j int
	fmt.Scanf("%d", &i)
	switch i % 2 {
	case 0:
		for j = i - 1; j > 2; {
			if IsPrime(j) {
				break
			}
			j = j - 2
		}
	case 1:

		for j = i; j > 2; {
			if IsPrime(j) {
				break
			}
			j = j - 2
		}
	}

	fmt.Printf("The Max Prime is %d ", j)

}

func IsPrime(i int) bool {
	for j := 3; j*j <= i; {
		if i%j == 0 {
			return false
		}
		j = j + 2
	}
	return true
}
