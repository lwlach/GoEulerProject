package main

import (
	"fmt"
	"strconv"

	"github.com/lwlach/GoEulerProject/common"
)

func main() {
	factorial100 := common.Factorial(100)
	fmt.Println(factorial100.String())
	var sum int
	for _, r := range factorial100.String() {
		n, _ := strconv.Atoi(string(r))
		sum += n
	}
	fmt.Println(sum)
}
