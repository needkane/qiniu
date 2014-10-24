package main

import (
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("/home/qboxtest/Desktop"))
	http.ListenAndServe(":3808", h)
}
