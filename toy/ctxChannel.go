package main

import "log"

func main() {
	ch := make(chan int, 1)
	ch <- 1
	//	close(ch)
	for {
		select {
		case <-ch:
			log.Println("case----")
		default:
			log.Println("default------")
		}

	}
}
