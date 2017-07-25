package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	var (
		lines []string
	)
	f, _ := os.Open("base64.txt")
	reader := bufio.NewReader(f)
	buffer := bytes.NewBuffer(make([]byte, 1024))
	for {
		part, prefix, err := reader.ReadLine()
		if err != nil {
			break
		}
		buffer.Write(part)

		if !prefix {
			res, err1 := http.Get(buffer.String())
			if err1 != nil || res.Status != "200" {
				lines = append(res.Header["X-Reqid"])
				buffer.Reset()
			}
			defer res.Body.Close()
		}
		if err != nil || io.EOF == err {
			break
		}
	}
	fout, _ := os.Create("result.txt")
	defer fout.Close()
	for _, elem := range lines {
		_, err := fout.WriteString(strings.TrimSpace(elem) + "\n")
		if err != nil {
			fmt.Println(err)
			break
		}
	}

}
