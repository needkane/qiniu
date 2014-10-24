package main

import "github.com/gographics/imagick/imagick"
import "fmt"

func main() {
	imagick.Initialize()
	defer imagick.Terminate()

	mw := imagick.NewMagickWand()
	defer mw.Destroy()
	//mw.ReadImage("/home/queniao/image/desktop.jpg")
	//mw.ReadImage("/home/queniao/image/7baa610e17872b81adb9daf331d720c8.jpg")
	mw.ReadImage("http://needkane.qiniudn.com/p.png")

	pw, _ := mw.GetImagePixelColor(1, 1)
	r := pw.GetRed()
	g := pw.GetGreen()
	b := pw.GetBlue()
	fmt.Println(r*65535, g*65535, b*65535)
}
