package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	p := math.Pow(2, 1000)
	var sum int
	numberStr := fmt.Sprintf("%f", p)
	for _, nStr := range numberStr {
		n, err := strconv.Atoi(string(nStr))
		if err != nil {
			break
		}
		sum += n
	}
	fmt.Println(sum)
}
