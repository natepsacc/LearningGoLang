package main

import (
	"fmt"
  "time"
)

func primeLoop(startNum int, maxNum int, respChannel chan []int)  {
	//startTime := time.Now()
//	max := 10000000
	primeSlice := []int{}
	for x := startNum; x < maxNum; x++ { 
		prime := true
		for y := 2; y*y <=x; y ++ {
			if x % y == 0 {
				prime = false
				break
			}
		}
		if prime == true {
			primeSlice = append(primeSlice, x)
		}
	}
	
	//elapsed := time.Since(startTime) //.Since, Nice!

	//fmt.Printf(fmt.Sprintf("identified %d primes in %s seconds", len(primeSlice), elapsed))

	respChannel <- primeSlice

}


func main() {
	startTime := time.Now()
	max := 10000000

	a_max := max / 2
	a_start := 2

	b_max := max 
	b_start := a_max + 1

	channel := make(chan []int) 

	go primeLoop(a_start, a_max, channel)
	go primeLoop(b_start, b_max, channel)

	resp := <- channel
	respB := <- channel


	resp = append(resp, respB...)
	elapsed := time.Since(startTime)
	fmt.Printf(fmt.Sprintf("identified %d primes in %s seconds", len(resp), elapsed))


}
