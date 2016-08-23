package main

import (
	"fmt"
	"log"
	"strconv"
	"strings"
)

func main() {
	is := []int{2, 6, 10, 12}
	var s string
	for _, v := range is {
		s = s + fmt.Sprintf(",00:00:%d,", v)
	}
	log.Println(strings.Trim(s, ","))
	s = "dasdsad"
	ss := strings.Split(s, ".")
	for _, v := range ss {
		log.Println(v)
	}

	mp := map[string]int{"pili-noded": 1, "pili-codecd": 12}

	for k, v := range mp {
		log.Println(k, "--------", v)
	}
	i, err := strconv.Atoi("214748364700")
	log.Println(i, err)
}
