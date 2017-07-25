package main

import (
	"encoding/base64"

	"fmt"
)

func main() {
	str := "aHR0cDovL2RuLXhzd2UucWJveC5tZS8xMDAzNzEzMzU+aW1hZ2VNb2dyMi90aHVtYm5haWwv"
	data, _ := base64.StdEncoding.DecodeString(str)
	fmt.Printf("%q\n", data)
}
