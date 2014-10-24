package main

import (
	"fmt"
	"github.com/gographics/imagick/imagick"
)

func main() {
	imagick.Initialize()
	defer imagick.Terminate()
	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	mw.ReadImage("http://needkane.qiniudn.com/r.png")

	height := int(mw.GetImageHeight())
	width := int(mw.GetImageWidth())
	println(height, width)
	pw := imagick.NewPixelWand()
	defer pw.Destroy()
	mw.GetImagePixelColor(2, 2)
	//fmt.Println("xxxx", pw)

	/*r := pw.GetRed()
	g := pw.GetGreen()
	b := pw.GetBlue()
	a := pw.GetAlpha()
	println("xxxxxxx", r, g, b, a)*/

	rgb := int(10)<<16 | int(11)<<8 | int(12)
	fmt.Printf("0x%06x", rgb)
}
