package main

import (
	"fmt"
	"net/http"
	"time"
)

func callback(w http.ResponseWriter, r *http.Request) {

	fmt.Println(r.RemoteAddr)
	time.Sleep(time.Second / 5)
	return
}

func main() {

	http.HandleFunc("/n", callback)
	panic(http.ListenAndServe(":8080", nil))
}

