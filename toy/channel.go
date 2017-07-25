package main

import (
	"fmt"
	"time"
)

func main() {
	go say("hello goroutine")
	say("just say")

	fmt.Println("------------------1")

	a := []int{7, 2, 8, -9, 4, 0}
	c := make(chan int)
	go sum(a[:len(a)/2], c)
	go sum(a[len(a)/2:], c)
	x, y := <-c, <-c

	fmt.Println(x, y, x+y)
	fmt.Println("------------------2")

	c2 := make(chan int, 2)
	c2 <- 1
	c2 <- 2
	fmt.Println(<-c2)
	fmt.Println(<-c2)
	fmt.Println("------------------3")

	c3 := make(chan int, 10)
	go fibonacci(cap(c3), c3)
	for i := range c3 {
		fmt.Println(i)
	}
	fmt.Println("------------------4")

	c4 := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c4)
		}
		quit <- 0
	}()
	fibonacci2(c4, quit)
}

func say(s string) {
	time.Sleep(100 * time.Millisecond)
	fmt.Println(s)
}

func sum(a []int, c chan int) {
	sum := 0
	for _, v := range a {
		sum += v
	}
	c <- sum
}
func fibonacci(n int, c chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func fibonacci2(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
			/*default:
			fmt.Println("Jane")*/
		}

	}
}

func name() {
	var chans = []chan uint32{srC, sgC, sbC, countC}

	for i := 0; i < len(chans); i++ {
		close(chans[i])
	}
	for i := range srC {
		srA += i
	}
	for i := range sgC {
		sgA += i
	}
	for i := range sbC {
		sbA += i
	}
	for i := range countC {
		countA += i
	}
	if countA == 0 {
		return
	}
}
