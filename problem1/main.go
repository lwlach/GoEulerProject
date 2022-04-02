package main

import "fmt"

func main() {
	var sum int
	for i := 0; i < 1000; i++ {
		if isMultiple(3, i) || isMultiple(5, i) {
			sum += i
		}
	}
	fmt.Println(sum)
}

func isMultiple(multiple, n int) bool {
	return n%multiple == 0
}
