package main

import (
	"fmt"

	"github.com/lwlach/GoEulerProject/common"
)

func main() {
	var sums = make(map[int]int)
	for i := 1; i < 10000; i++ {
		sums[i] = common.SumOfDivisors(i)
	}
	var sum int
	for k, v := range sums {
		for k2, v2 := range sums {
			if k == v2 && v == k2 && k != v {
				sum += k + v
				delete(sums, k)
			}
		}
	}
	fmt.Println(sum)
}
