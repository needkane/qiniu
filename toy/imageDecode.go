package main

import (
	"fmt"
	"image"
	"os"

	_ "code.google.com/p/go.image/bmp"
	_ "code.google.com/p/vp8-go/webp"
	_ "image/gif"
	_ "image/png"
)

func main() {
	imgFile, err := os.Open("/home/qboxtest/Desktop/597Ave")
	m, _, err := image.Decode(imgFile)
	fmt.Println(m, err)

}
