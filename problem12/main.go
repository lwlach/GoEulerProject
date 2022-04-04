package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println(time.Now())
	var (
		divisors, n   int
		triangleIndex = 1
	)
outer:
	for {
		divisors = 0
		triangleIndex++
		n = triangleNumber(triangleIndex)
		for i := 1; i <= n; i++ {
			if n%i == 0 {
				divisors++
				if divisors > 500 {
					break outer
				}
			}
		}
	}
	fmt.Println(n)
	fmt.Println(time.Now())
}

func triangleNumber(index int) int {
	var n int
	for i := 1; i <= index; i++ {
		n += i
	}
	return n
}
