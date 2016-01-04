package main

import (
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("/home/qboxtest/Downloads"))
	http.ListenAndServe(":1935", h)
}
