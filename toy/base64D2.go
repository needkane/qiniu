package main

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

func main() {
	var (
		lines []string
	)
	f, _ := os.Open("/home/qboxtest/Desktop/a9")
	reader := bufio.NewReader(f)
	buffer := bytes.NewBuffer(make([]byte, 1024))
	for {
		part, prefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		buffer.Write(part)

		if !prefix {
			data, _ := base64.StdEncoding.DecodeString(buffer.String())
			lines = append(lines, string(data))
			buffer.Reset()
		}
		if err != nil || io.EOF == err {
			break
		}
	}

	fout, _ := os.Create("base64.txt")
	defer fout.Close()
	for _, elem := range lines {
		_, err := fout.WriteString(strings.TrimSpace(elem) + "\n")
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}
