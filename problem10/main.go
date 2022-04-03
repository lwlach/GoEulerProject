package main

import (
	"ProjectEuler/common"
	"fmt"
)

func main() {
	var sum int
	for i := 2; i < 2000000; i++ {
		if common.IsPrime(i) {
			sum += i
		}
	}
	fmt.Println(sum)
}
