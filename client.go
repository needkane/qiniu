package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {

	client := http.Client{Transport: http.DefaultTransport}

	{
		req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/n", nil)
		go func() {
			time.Sleep(time.Second / 10)
			client.Transport.(requestCanceler).CancelRequest(req)
		}()
		_, err := client.Do(req)
		fmt.Println("err:", err)
	}

	{
		req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/n", nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("err2:", err)
		} else {
			fmt.Println("resp2 code:", resp.StatusCode)
			resp.Body.Close()
		}
	}

	{
		req, _ := http.NewRequest("GET", "http://127.0.0.1:8080/n", nil)
		resp, err := client.Do(req)
		if err != nil {
			fmt.Println("err3:", err)
		} else {
			fmt.Println("resp3 code:", resp.StatusCode)
			resp.Body.Close()
		}
	}
}

type requestCanceler interface {
	CancelRequest(*http.Request)
}

