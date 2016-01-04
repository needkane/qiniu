package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"
)

type errorRet struct {
	Error string `json:"error"`
}

func main() {
	Listen()
}

func Listen() {

	//define handler
	http.HandleFunc("/uop", serveUfop)

	//bind and listen

	ufopServer := &http.Server{
		Addr: "0.0.0.0:6666",
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
	}
	defer req.Body.Close()
	ufopReqData, err := ioutil.ReadAll(req.Body)
	if err != nil {
		WriteJsonError(w, 500, "Ufop: read req.Body failed")
	}
	log.Println(string(ufopReqData))
}

func WriteJsonError(w http.ResponseWriter, code int, content string) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	msg, _ := json.Marshal(errorRet{content})
	w.Write(msg)
}
