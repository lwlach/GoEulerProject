package main

import "fmt"

func main() {
	var (
		fib       = []int{1, 2}
		fibNumber = fib[0] + fib[1]
		sum       = 2
	)
	for fibNumber <= 4000000 {
		fib[0] = fib[1]
		fib[1] = fibNumber
		if fibNumber%2 == 0 {
			sum += fibNumber
		}
		fibNumber = fib[0] + fib[1]
	}
	fmt.Println(sum)
}
