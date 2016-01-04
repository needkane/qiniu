package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

func main() {

	f, _ := os.Open("routine.txt")
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)

	reader := bufio.NewReader(f)
	for {
		line, err1 := reader.ReadString('\n')
		if err1 == io.EOF {
			break
		}
		go fmt.Println("----", line)
	}
	defer f.Close()
}
