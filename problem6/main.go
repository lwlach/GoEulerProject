package main

import "fmt"

func main() {
	const N = 100
	var sumSquares, sum int
	for i := 1; i <= N; i++ {
		sumSquares += i * i
		sum += i
	}
	result := (sum * sum) - sumSquares
	fmt.Println(result)
}
