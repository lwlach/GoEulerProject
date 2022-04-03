package main

import (
	"ProjectEuler/common"
	"fmt"
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
