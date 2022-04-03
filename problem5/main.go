package main

import "fmt"

func main() {
	var (
		isAllEven bool
		dividend  = 20
	)
main:
	for !isAllEven {
		for i := 1; i <= 20; i++ {
			if !isEvenDivided(dividend, i) {
				dividend++
				continue main
			}
		}
		isAllEven = true
	}
	fmt.Println(dividend)
}

func isEvenDivided(dividend, divisor int) bool {
	return dividend%divisor == 0
}
