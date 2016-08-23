package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func main() {
	f, err := os.Open("nodelist")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	reader := bufio.NewReader(f)
	var s1 []string
	var s0 string
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		s0 += strings.TrimSuffix(line[2:len(line)-2], "\"")
	}
	f, err = os.Open("nodelist2")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	reader = bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Println(err)
		}
		s1 = append(s1, strings.TrimSuffix(line, "\n"))
	}
	var s2 string
	for _, i := range s1 {
		if !strings.Contains(s0, i) {
			s2 += i + ","
		}
	}
	fmt.Println(s2)
}
