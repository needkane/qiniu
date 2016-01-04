package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
)

func f(from string) {
	for i := 0; i < 10; i++ {
		fmt.Println(from, "...", i)
	}
}

func main() {
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)
	fmt.Println("PROCS:", NCPU)
	/*go f("direct")
	go f("goroutine")
	go func(msg string) {
		for i := 0; i < 10; i++ {
			fmt.Println("internal " + msg)
		}

	}("going")*/
	f, _ := os.Open("routine.txt")

	reader := bufio.NewReader(f)
	for {
		line, err1 := reader.ReadString('\n')
		if err1 == io.EOF {
			break
		}
		go fmt.Println("----", line)
	}
	var input string
	fmt.Scanln(&input)
	fmt.Println("done by " + input)
}
