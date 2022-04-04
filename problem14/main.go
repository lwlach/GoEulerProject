package main

import "fmt"

func main() {
	var largestChain, n int
	for i := 999999; i > 1; i-- {
		chain := collatzSequence(i)
		if chain > largestChain {
			largestChain = chain
			n = i
		}
	}
	fmt.Println(n)
}

func collatzSequence(n int) int {
	var count int
	for n != 1 {
		if n%2 == 0 {
			n = n / 2
		} else {
			n = (3 * n) + 1
		}
		count++
	}
	return count + 1
}
