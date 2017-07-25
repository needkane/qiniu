package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	resp, err := http.Get("http://127.0.0.1:3808/sex.mp4")
	if err != nil {
		log.Println("err: ", err)
		return
	}
	defer resp.Body().Close()
	os.Create("/home/qboxtest/Downloads/a")
}
