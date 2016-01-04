package main

import (
	"net/http"
	"os"
	"runtime"
	"strconv"
	//"strings"
)

func main() {
	runtime.GOMAXPROCS(8)
	f, _ := os.Open("/home/qboxtest/Desktop/IFD0NOoriIFD1HasOri")
	defer f.Close()

	//for i := 0; i < 10; i++ {

	f.Seek(0, 0)
	//req, err := http.NewRequest("POST", "http://192.168.200.161:29102/op?fsize=199069&cmd=imageMogr2/thumbnail/200x/"+strconv.Itoa(i), f)
	resp, _ := http.Post("http://192.168.200.161:29102/op?fsize=199069&cmd=imageMogr2/thumbnail/200x/"+strconv.Itoa(1), "image/jpeg", f)
	defer resp.Body.Close()
	//}

}
