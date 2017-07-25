package main

import (
	"fmt"
	"syscall"
)

func main() {
	var i int
	fmt.Scanf("%d", &i)
	if i == 0 {
		fmt.Println(syscall.Getpid())
		syscall.Kill(syscall.Getpid(), syscall.SIGHUP)
	}
	fmt.Println("-----------")
}
