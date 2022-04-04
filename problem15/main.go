package main

import (
	"fmt"
	"math/big"
)

func main() {
	const N = 20
	var factorial = big.NewInt(1)
	for i := 2; i <= N; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}
	var doubleFactorial = big.NewInt(1)
	for i := 2; i <= 2*N; i++ {
		doubleFactorial.Mul(doubleFactorial, big.NewInt(int64(i)))
	}
	paths := doubleFactorial.Div(doubleFactorial, factorial.Mul(factorial, factorial))
	fmt.Println(paths.String())
}
