package main

import (
	"fmt"
	"os"
	"path"
)

func main() {
	a := path.Join("http://", "a.com", "/dsds")
	fmt.Println(a)
	ss := []string{"qiniu", "127.0.0.1"}
	fmt.Println(ss[1:])
	Pa := os.Getenv("GOROOT")
	Pa = path.Join(Pa, "pili-zeus")
	fmt.Println(Pa)
}
