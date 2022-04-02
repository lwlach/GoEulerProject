package main

import (
	"fmt"
)

func main() {
	fmt.Println(largestPalindrome())
}

func largestPalindrome() int {
	var palindromes []int
	for a := 999; a >= 100; a-- {
		for b := 999; b >= 100; b-- {
			product := a * b
			strProduct := fmt.Sprintf("%d", product)
			if isPalindrome(strProduct) && max(palindromes) < product {
				//fmt.Println(a)
				//fmt.Println(b)
				palindromes = append(palindromes, product)
			}
		}
	}
	return palindromes[len(palindromes)-1]
}

func max(s []int) int {
	if len(s) == 0 {
		return -1
	}
	var m = s[0]
	for _, i := range s {
		if i > m {
			m = i
		}
	}
	return m
}

func isPalindrome(s string) bool {
	for i, r := range s {
		if i == len(s)/2 {
			return true
		}
		reversedIndex := len(s) - i - 1
		if r != rune(s[reversedIndex]) {
			return false
		}
	}
	return true //never gets here
}
