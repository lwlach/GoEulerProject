package main

import (
	"fmt"
	"math"
)

func main() {
	//fmt.Println(highestPrimeFactor(13195))
	fmt.Println(highestPrimeFactor(600851475143))
}

func highestPrimeFactor(n int) int {
	var (
		maxPrime int
		square   = int(math.Sqrt(float64(n)))
	)
	for i := square; i >= 2; i-- {
		if isPrime(i) && n%i == 0 {
			maxPrime = i
			break
		}
	}
	if maxPrime == 0 {
		return 1
	}
	return maxPrime
}

func isPrime(n int) bool {
	for i := 2; i < n; i++ {
		if n%i == 0 {
			return false
		}
	}
	return true
}
