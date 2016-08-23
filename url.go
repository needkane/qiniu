package main

import (
	"fmt"
	"log"
	"net/url"
	"strings"
)

func main() {
	s := "http://source.com/aa%3D"
	u, _ := url.Parse(s)
	fmt.Println(u.String())
	s1 := "1232131"
	fmt.Println(delScheme(s))
	log.Println(delScheme(s1))
}

func delScheme(rawurl string) string {
	if strings.HasPrefix(rawurl, "http://") {
		return rawurl[len("http://"):]
	}
	return rawurl
}
