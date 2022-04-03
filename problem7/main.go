package main

import (
	"fmt"

	"github.com/lwlach/GoEulerProject/common"
)

func main() {
	var (
		count int
		i     = 1
	)
	for count < 10001 {
		i++
		if common.IsPrime(i) {
			count++
		}
	}
	fmt.Println(i)
}
