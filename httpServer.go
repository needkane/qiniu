package main

import (
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
)

func main() {

	//define handler
	http.HandleFunc("/uop", serveUfop)

	//bind and listen

	ufopServer := &http.Server{
		Addr: "127.0.0.1:29109",
	}

	listenErr := ufopServer.ListenAndServe()
	if listenErr != nil {
		log.Println(listenErr)
	}
}

func serveUfop(w http.ResponseWriter, req *http.Request) {
	//check method
	if req.Method != "POST" {
		WriteJsonError(w, 405, "method not allowed")
		return
	}
	defer req.Body.Close()
	_, err := ioutil.ReadAll(req.Body)
	if err != nil {
		WriteJsonError(w, 500, "Ufop: read req.Body failed")
		return
	}

	ReplyImageStream(w)
}

func WriteJsonError(w http.ResponseWriter, code int, content string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)

	w.Write([]byte("msg"))
}

func ReplyImageStream(w http.ResponseWriter) {
	f, _ := os.Open("/home/qboxtest/qbox/devtools/bin/a.m3u8")
	defer f.Close()
	fi, _ := f.Stat()

	fsize := fi.Size()

	w.Header().Set("Content-Type", "")
	w.Header().Set("Content-Length", strconv.FormatInt(fsize, 10))
	w.WriteHeader(200)
	io.CopyN(w, f, fsize)
}
