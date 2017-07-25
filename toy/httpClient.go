package main

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"
)

func main() {
	log.Println("--------")
	client := http.Client{Timeout: 2 * time.Second}
	req, _ := http.NewRequest("GET", "http://127.0.0.1:8881/listnodes", nil)
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	bs, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	log.Println(string(bs))
}
