package main

import (
	"fmt"

	"github.com/lwlach/GoEulerProject/common"
)

func main() {
	var maxPrimes, maxAB int
	for a := -999; a <= 999; a++ {
		for b := -1000; b <= 1000; b++ {
			n := countPrimes(a, b)
			if n > maxPrimes {
				fmt.Println("a", a, "b", b, "n", n)
				maxPrimes = n
				maxAB = a * b
			}
		}
	}
	fmt.Println(maxPrimes, maxAB)
}

func countPrimes(a, b int) int {
	var n int
	for common.IsPrime(n*n + a*n + b) {
		n++
	}
	return n
}
