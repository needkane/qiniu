package main

import (
	"fmt"
	//"io"
	"io/ioutil"
	"os"
)

func main() {
	f, _ := os.Open("/home/qboxtest/Desktop/NoOri.jpg")

	defer f.Close()
	fi, _ := f.Stat()
	fmt.Println("xxx", fi)

	var bytes = make([]byte, 1000)
	fmt.Println("zzzzzzzzxxxxxxxxz", bytes)
	f.Read(bytes)
	fmt.Println("zzzzzzzzz", bytes)
	bytes, err := ioutil.ReadAll(f)
	fmt.Println("yyyyyyyyyyy", bytes, f, err)
}
