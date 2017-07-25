package main

import (
	"fmt"
	"image"
	"os"

	_ "image/jpeg"
)

func main() {
	f, _ := os.Open("/home/qboxtest/Desktop/gvs")
	m, format, err := image.Decode(f)
	if err != nil {
		fmt.Println("abnormal image", format)
	} else {
		fmt.Println(m.ColorModel())
	}

}
