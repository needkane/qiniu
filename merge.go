package main

import (
	"io"
	"log"
	"os"
)

func main() {
	fs, err := os.Create("c.txt")
	if err != nil {
		log.Println("create failed: ", err)
	}
	f, err := os.Open("a.txt")
	if err != nil {
		log.Println("open a.txt failed: ", err)
	}
	f1, err := os.Open("b.txt")
	if err != nil {
		log.Println("open b.txt failed: ", err)
	}
	//tr := io.NewReader(f1)
	//	tw := timeio.NewWriter(fs)
	n, err := io.Copy(fs, f1)
	log.Println(n, "---------", err)

	//tr = timeio.NewReader(f)
	//	tw = timeio.NewWriter(fs)
	n, err = io.Copy(fs, f)
	log.Println(n, "---------", err)
}
