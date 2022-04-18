package main

import "fmt"

func main() {
	var maxCycle int
	for n := 2; n < 1000; n++ {
		nCp := makeCoPrimeToTen(n)
		if nCp != 1 {
			cycle := getCycleLength(nCp)
			if cycle > maxCycle {
				maxCycle = cycle
				fmt.Println(fmt.Sprintf("1/%d has a %d digit recurring cycle", n, cycle))
			}
		}
	}
}

// Function to remove all factors that are multiples of 2 and multiples of 5 to make it co-prime to 10.
func makeCoPrimeToTen(n int) int {
	for n%2 == 0 {
		n = n / 2
	}
	for n%5 == 0 {
		n = n / 5
	}
	return n
}

func getCycleLength(n int) int {
	var (
		modVal  = 10 % n
		currMod = modVal
		k       = 1
	)
	for {
		if currMod == 1 {
			return k
		}
		currMod = (currMod * modVal) % n
		k++
	}
}
