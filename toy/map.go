package main

import "log"

func main() {
	mp := map[string]string{"t1": "v1", "t2": "v2"}
	i, v := mp["t3"]
	log.Println(i, "-------", v)
	for i, v := range mp {
		log.Println(i, "---------", v)
	}

	var mp2 = make(map[string]int, 0)
	mp2["a"] = 1
	mp2["n"] = 2
	log.Println(mp2)
}
