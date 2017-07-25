package main

import (
	"fmt"
	"image"
	"os"
	"runtime"
)

func main() {
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("OSX")
	case "linux":
		fmt.Println("Linux")
	default:
		//freebsd windows openbsd plan9
		fmt.Printf("os is %s", os)
	}
	var i float64
	println(i)

	f, _ := os.Open("/home/qboxtest/cantRead")
	im, _, err := image.DecodeConfig(f)
	fmt.Println(im.Width, im.Height, err)

	a := "abcd"
	fmt.Println(a[len(a)-1 : len(a)])
}
