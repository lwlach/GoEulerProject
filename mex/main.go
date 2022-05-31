package main

import (
	"fmt"
	"sort"
	"time"
)

func main() {
	var arr []int
	for i := 100000000; i >= 0; i-- {
		if i == 657832 {
			continue
		}
		arr = append(arr, i)
	}
	fmt.Println(time.Now())
	mex := findMex(arr)
	fmt.Println(mex)
	fmt.Println(time.Now())
}

func findMex(arr []int) int {
	notMex := make(map[int]bool)
	for _, n := range arr {
		notMex[n] = true
	}
	var mex int
	for i := 0; i <= len(arr); i++ {
		if !notMex[i] {
			mex = i
			break
		}
	}
	return mex
}

// faster
func findMex2(arr []int) int {
	sort.Ints(arr)
	for i, n := range arr {
		if i != n {
			return i
		}
	}
	return 0
}
