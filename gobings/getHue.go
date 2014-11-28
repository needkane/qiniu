package main

import (
	"fmt"
	"github.com/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	err := mw.ReadImage("/home/qboxtest/Desktop/111.jpg")
	fmt.Println(err)
	height := int(mw.GetImageHeight())
	width := int(mw.GetImageWidth())
	println(height, width)
	//pw := imagick.NewPixelWand()
	for i := 0; i < height*width; i++ {
		x := i % width
		y := i / width

		pw, _ := mw.GetImagePixelColor(x, y)
		defer pw.Destroy()
		r := uint32(pw.GetRed() * 65535)
		g := uint32(pw.GetGreen() * 65535)
		b := uint32(pw.GetBlue() * 65535)
		a := uint32(pw.GetAlpha() * 65535)
		println("xxxxxxx", r, g, b, a)

	}
	rgb := int(10)<<16 | int(11)<<8 | int(12)
	fmt.Printf("0x%06x", rgb)
}
