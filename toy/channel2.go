package main

import (
	"fmt"
)

func main() {
	c := make(chan int)
	go func() {
		vs := []int{1, 3, 5, 7}
		for _, v := range vs {
			c <- v
		}
		close(c)
	}()

	for v := range c {
		fmt.Println(v)
	}
	fmt.Println(<-c)
}
