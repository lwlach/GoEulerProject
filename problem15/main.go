package main

import (
	"fmt"

	"github.com/lwlach/GoEulerProject/common"
)

func main() {
	const N = 20
	var factorial = common.Factorial(N)
	var doubleFactorial = common.Factorial(2 * N)
	paths := doubleFactorial.Div(doubleFactorial, factorial.Mul(factorial, factorial))
	fmt.Println(paths.String())
}
