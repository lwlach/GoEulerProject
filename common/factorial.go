package common

import "math/big"

func Factorial(n int) *big.Int {
	var factorial = big.NewInt(1)
	for i := 2; i <= n; i++ {
		factorial.Mul(factorial, big.NewInt(int64(i)))
	}
	return factorial
}
