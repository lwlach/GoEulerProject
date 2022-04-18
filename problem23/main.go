package main

import (
	"fmt"

	"github.com/lwlach/GoEulerProject/common"
)

func main() {
	const threshHold = 28123
	var (
		abundants     []int
		sumsAbundants = make(map[int]bool)
	)

	for i := 1; i <= threshHold; i++ {
		if common.SumOfDivisors(i) > i {
			abundants = append(abundants, i)
		}
	}
	for _, a := range abundants {
		for _, a2 := range abundants {
			sumsAbundants[a+a2] = true
		}
	}

	var sumNonAbundant int
	for i := 1; i <= threshHold; i++ {
		if !sumsAbundants[i] {
			sumNonAbundant += i
		}
	}
	fmt.Println(sumNonAbundant)
}
