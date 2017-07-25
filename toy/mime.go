package main

import (
	"fmt"
)

var mimes = map[string]string{
	".%":   "application/x-trash",
	".123": "application/vnd.lotus-1-2-3",
	".323": "text/h323",
}

func main() {
	fmt.Println("----------", ContentType(".application/vnd.apple.mpegurl"))
}

func ContentType(ext string) string {
	return mimes[ext]
}
