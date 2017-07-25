package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"qbox.us/net/httputil"
)

type StreamStatusResponse struct {
	Status string `json:"status"`
}

var (
	checked  = "rtmp://10.200.20.31:1835/result/strR"
	notFound = "rtmp://127.0.0.1/9xiu/str2"
	failed   = "rtmp://127.0.0.1/9xiu/str3"
)

type Result struct {
	Status string `json"status"`
}

type Input struct {
	Urls []string `json"urls"`
}

func main() {

	//define handler
	http.HandleFunc("/stream/status", serveUfop)

	//bind and listen

	ufopServer := &http.Server{
		Addr: "0.0.0.0:2803",
	}

	listenErr := ufopServer.ListenAndServe()
	if listenErr != nil {
		log.Println(listenErr)
	}
}

func serveUfop(w http.ResponseWriter, req *http.Request) {
	//check method
	var (
		SourceBytes []byte
		err         error
	)
	if req.Method != "POST" {
		err = httputil.NewError(405, "method not allowed")
		httputil.Error(w, err)
		return

	}
	defer req.Body.Close()
	SourceBytes, err = ioutil.ReadAll(req.Body)
	if err != nil {
		httputil.Error(w, err)
		return

	}

	var res Input
	err = json.Unmarshal(SourceBytes, &res)
	if err != nil {
		httputil.Error(w, err)
		return

	}

	switch res.Urls[0] {
	case checked:
		WriteResult(w, 200, "checked")
	case notFound:
		WriteResult(w, 200, "not found")
	case failed:
		WriteResult(w, 200, "failed")

	}
}

func WriteResult(w http.ResponseWriter, code int, content string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	msg, _ := json.Marshal(Result{content})
	w.Write(msg)
}

func ReplyImageStream(w http.ResponseWriter, code int, content string) {

	w.Header().Set("Content-Type", "")
	w.WriteHeader(code)
	w.Write([]byte(content))
}
