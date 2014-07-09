package main

import (
	"./stacker"
	"fmt"
)

func main() {

	var hayStack stacker.Stack
	hayStack.Push("hey")
	hayStack.Push("15")
	hayStack.Push([]string{"pin", "clip"})
	for {
		item, err := hayStack.Pop()
		if err != nil {
			break
		}
		fmt.Println(item)
	}
}
