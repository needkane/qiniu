package main

import (
	"fmt"
	"math"
	"sync/atomic"
)

func main() {
	var x, cur uint64
	cur = math.MaxUint64
	fmt.Println(cur, cur+1)
	for i := 0; i < 10; i++ {
		x = atomic.AddUint64(&cur, 1)
	}
	fmt.Println(x, cur)
}
