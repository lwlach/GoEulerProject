package main

import (
	"fmt"
	"sort"
	"strings"
)

var permutations []string

func main() {
	var (
		digits    = []int{0, 1, 2}
		digitsStr []string
	)
	for _, d := range digits {
		digitsStr = append(digitsStr, fmt.Sprintf("%d", d))
	}
	permutation(digitsStr, len(digits))
	sort.Strings(permutations)
	fmt.Println(permutations[1000000-1])
}

func permutation(arr []string, size int) {
	if size == 1 {
		fmt.Println(arr)
		permutations = append(permutations, strings.Join(arr, ""))
		return
	}
	for i := 0; i < size; i++ {
		permutation(arr, size-1)
		last := size - 1
		if size%2 == 1 {
			arr[0], arr[last] = arr[last], arr[0]
		} else {
			arr[i], arr[last] = arr[last], arr[i]
		}
	}
}
