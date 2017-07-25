package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	fa, err := os.Open("aa")
	if err != nil {
		log.Println(err)
	}

	fc, err := os.Open("c2")
	if err != nil {
		log.Println(err)
	}
	var aa string
	var bb = []string{}
	reader := bufio.NewReader(fa)
	for {
		line, err1 := reader.ReadString('\n')
		if err1 == io.EOF {
			break
		}
		aa = aa + line
	}
	reader2 := bufio.NewReader(fc)
	for {
		line, err1 := reader2.ReadString('\n')
		if err1 == io.EOF {
			break
		}
		bb = append(bb, line)
	}
	//	log.Println(aa)
	//	log.Println(bb)
	var cc = []string{}
	for _, v := range bb {
		if !strings.Contains(aa, v) {
			cc = append(cc, v)
		}
	}
	log.Println(cc)
}
