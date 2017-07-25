package main

import (
	"fmt"
	"time"
)

func main() {
	/*	stopPoint := make(chan bool, 1)
		log.Println("--------")
		go func() {
			for {
				select {
				case <-stopPoint:
					log.Println("==========point")
				case <-time.After(3 * time.Second):
					log.Println("------------after")
				}

			}
		}()

		go func() {
			time.Sleep(10 * time.Second)
			stopPoint <- true
		}()
		tn := time.Now()
		time.Sleep(13 * time.Second)
		log.Println(int(time.Since(tn) / time.Millisecond))
	*/
	p := fmt.Println
	t := time.Now()
	p(t.Format("2006-01-02/15:04"))
	fmt.Println(t.Year(), "-", int(t.Month()), "-", t.Day(), "/", t.Hour(), ":", t.Minute())
	p(time.Now().Format("2016-07-28/12:04"))
}
