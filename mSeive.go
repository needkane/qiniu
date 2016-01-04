package main

import (
	"fmt"
	"math"
)

/*// Send the sequence 2, 3, 4, ... to channel 'ch'.
func Generate(ch chan<- int) {
	for i := 2; i < 110; i++ {
		ch <- i // Send 'i' to channel 'ch'.
	}
}

// Copy the values from channel 'in' to channel 'out',
// removing those divisible by 'prime'.
func Filter(in <-chan int, out chan<- int, prime int) {
	for {
		i := <-in // Receive value from 'in'.
		if i%prime != 0 {
			out <- i // Send 'i' to 'out'.
		}
	}
}*/

// The prime sieve: Daisy-chain Filter processes.
func main() {
var  i, j int
   for i:=2; i<100; i++{
  if(i%2) prime=false
}
  else prime=true;
   for(i=3; i<=sqrt(N); i+=2)
   {   if(prime)
       for(j=i+i; j<N; j+=i)
            prime=false;
   }
   for(i=2; i<100; i++)//由于输出将占用太多io时间，所以只输出2-100内的素数。可以把100改为N
    if( prime )
         printf("%d ",i);
   return 0;
}
