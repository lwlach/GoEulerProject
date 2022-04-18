package main

import (
	"fmt"
	"math/big"
)

func main() {
	var (
		index     = 1000
		fibDigits int
	)
	for fibDigits < 1000 {
		fibDigits = len(fibonacci(index).String())
		index++
	}
	fmt.Println(index)
}

func fibonacci(n int) *big.Int {
	var (
		count = 2
		fib   = []*big.Int{big.NewInt(1), big.NewInt(1)}
	)
	for n > count {
		nextFib := new(big.Int)
		nextFib.Add(fib[0], fib[1])
		fib[0] = fib[1]
		fib[1] = nextFib
		count++
	}
	return fib[1]
}
