package main

import (
	"fmt"
	"sort"
)

func main() {
	var (
		// 2
		p = []int{1, 4, 1}
		s = []int{1, 5, 1}

		// 3
		//p = []int{4, 4, 2, 4}
		//s = []int{5, 5, 2, 5}

		// 2
		//p = []int{2, 3, 4, 2}
		//s = []int{2, 5, 7, 2}
	)
	fmt.Println(Solution(p, s))
}

func Solution(P []int, S []int) int {
	var people int
	for _, n := range P {
		people += n
	}
	var cars int
	sort.Ints(S)
	for i := len(S) - 1; i >= 0; i-- {
		if people > 0 {
			people -= S[i]
			cars++
		} else {
			break
		}
	}
	return cars
}
