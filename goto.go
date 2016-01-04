package main

import "fmt"

func main() {
	var ss = []string{"kane", "use", "goto", "instead", "of", "for"}
	myGoto(ss)
}

func myGoto(ss []string) {
	fmt.Println(ss[0])
	if len(ss) > 1 {
		myGoto(ss[1:])
	}
}
