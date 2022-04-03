package main

import "fmt"

func main() {
outer:
	for c := 3; c < 999; c++ {
		for b := 2; b < c; b++ {
			for a := 1; a < b; a++ {
				if a+b+c == 1000 {
					cs := c * c
					bs := b * b
					as := a * a
					if as+bs == cs {
						fmt.Println(a, b, c)
						fmt.Println(a * b * c)
						break outer
					}
				}
			}
		}
	}
}
