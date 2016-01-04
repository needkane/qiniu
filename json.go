package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	a := "{\"error\":\"thumbnail's arg is out of range: 621213455678x\"}"
	m := make(map[string]interface{})
	json.Unmarshal([]byte(a), &m)

	fmt.Println("xxxx", m["error"])
}
