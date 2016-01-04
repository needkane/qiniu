package main

import (
	"fmt"
	"regexp"
)

func main() {
	var (
		matched bool
		err     error
	)
	s := "300x300+0-10"
	if matched, err = regexp.MatchString(`[0-9]+[x][0-9]+[+][0-9]+[+][0-9]+`, s); matched {
		fmt.Println("1")
	} else if matched, err = regexp.MatchString(`[0-9]+[x][0-9]+[+][0-9]+[-][0-9]+`, s); matched {
		fmt.Println("2")
	} else if matched, err = regexp.MatchString(`[0-9]+[x][0-9]+[-][0-9]+[-][0-9]+`, s); matched {
		fmt.Println("3")
	} else if matched, err = regexp.MatchString(`[0-9]+[x][0-9]+[-][0-9]+[+][0-9]+`, s); matched {
		fmt.Println("4")
	} else {
		fmt.Println(err)
	}

	/*reg := regexp.MustCompile(`[0-9]+`)
	fmt.Println(reg.Find("500x500+12+13"))*/

	matched,err=regexp.Match("^faceCrop/[0-9]+[x][0-9]+",[]byte("faceCrop/50f0x500"))
	fmt.Println("------",matched,err)
	
}
