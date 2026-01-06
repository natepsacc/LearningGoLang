package main

import (
	"fmt"
	"time"
)

func main() {
	startTime := time.Now()
	max := 100000
	primeSlice := []int{}
	for x := 0; x < max; x++ { 
		prime := true
		for y := 2; y < x -1; y ++ {
			if x % y == 0 {
				prime = false
				break
			}
		}
		if prime == true {
			primeSlice = append(primeSlice, x)
		}
	}
	
	elapsed := time.Since(startTime) //.Since, Nice!

	fmt.Printf(fmt.Sprintf("identified %d primes in %s seconds", len(primeSlice), elapsed))

}
