package main

import (
	"bufio"
	//"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	f, _ := os.Open("base64")
	reader := bufio.NewReader(f)
	for {
		line, err := reader.ReadString('\n')
		if err != nil || io.EOF == err {
			break
		}
		ss := strings.Split(line, " ")
		line = ss[1]
		if strings.Contains(line, "%3D") {
			line = strings.Replace(line, "%3D", "", -1)
		}
		if strings.Contains(line, "_") {
			line = strings.Replace(line, "_", "+", 1)
		}
		bytes, err := base64.StdEncoding.DecodeString(line)
		//fmt.Printf("%q\n", bytes)
		str := string(bytes)
		if strings.Contains(str, ">") {
			str = strings.Replace(str, ">", "?", 1)
		}
		fmt.Println(ss[0] + "," + str)
	}

}
