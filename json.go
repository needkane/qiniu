package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type TransCode struct {
	Src   string  `json:"src"`
	Dests []*Dest `json:"dests"`
}

type Dest struct {
	url      string `json:"url"`
	codecCmd string `json:"codecCmd"`
}

type Output struct {
}

func main() {
	a := "{\"error\":\"thumbnail's arg is out of range: 621213455678x\"}"
	m := make(map[string]interface{})
	json.Unmarshal([]byte(a), &m)

	fmt.Println("xxxx", m["error"])
	f, _ := os.Open("json.txt")
	bytes, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("read failed:", err)
	}
	var ret TransCode
	err = json.Unmarshal(bytes, &ret)
	if err != nil {
		fmt.Println("json failed:", err)
	} else {
		fmt.Println(ret)
	}
}
