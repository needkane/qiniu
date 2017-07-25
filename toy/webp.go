package main

import(
	"image"
	_ "golang.org/x/image/webp"
	"fmt"
	"os"
	_ "image/gif"
)

func main(){
	//f,_:=os.Open("/home/qboxtest/Downloads/youlin.webp")
	f,_:=os.Open("goWrongGif")
	_, format, err := image.DecodeConfig(f)
	fmt.Println(err,"-----",format)
}
