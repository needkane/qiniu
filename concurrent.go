package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"runtime"
	"strconv"
	"time"
)

func printHelp() {
	fmt.Println("usage: cmd [input file] [serve url] [concurrent times]")
}
func main() {
	if len(os.Args) < 4 {
		printHelp()
		os.Exit(-1)
	}
	var input = os.Args[1]
	var url = os.Args[2]
	con, err := strconv.Atoi(os.Args[3])
	if err != nil {
		fmt.Println("[concurrent times] should int type")
	}
	log.Println(input, "------", url)
	NCPU := runtime.NumCPU()
	runtime.GOMAXPROCS(NCPU)
	fmt.Println("PROCS:", NCPU)
	log.Println("start")
	for {
		for i := 0; i < con; i++ {
			go func() {
				//"/home/qboxtest/Downloads/towebp.jpg"
				f, err := os.Open(input)
				if err != nil {
					fmt.Println("open faile ", err)
				}
				defer f.Close()
				//"http://127.0.0.1:14102/op?fsize=4350755&cmd=imageMogr2/thumbnail/500x/format/webp"
				resp, err := http.Post(url, "image/jpg", f)
				if err != nil {
					fmt.Println("post failed ", err)
				}
				defer resp.Body.Close()
				body, err := ioutil.ReadAll(resp.Body)
				if err != nil {
					fmt.Println("read failed ", err)
				}
				fmt.Println(strconv.Itoa(i) + "-------" + string(body[:20]))

			}()
		}
		<-time.After(time.Second * 5)
	}

	fmt.Println("done by " + input)
}
