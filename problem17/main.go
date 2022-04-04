package main

import (
	"fmt"
	"strings"
)

func main() {
	var sum int
	for i := 1; i <= 1000; i++ {
		number := numToWord(i)
		number = strings.ReplaceAll(number, " ", "")
		number = strings.ReplaceAll(number, "-", "")
		sum += len(number)
	}
	fmt.Println(sum)
}

var zeroToNineteen = []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine", "ten", "eleven", "twelve", "thirteen", "fourteen", "fifteen", "sixteen", "seventeen", "eighteen", "nineteen"}

func numToWord(n int) string {
	if n < 20 {
		return zeroToNineteen[n]
	} else if n < 100 {
		return underHundred(n)
	} else if n < 1000 {
		return underThousand(n)
	} else if n < 1000000 {
		return underMillion(n)
	}
	return "invalid input"
}

func underMillion(n int) string {
	var thousand = "thousand"
	r := n % 1000
	d := n / 1000
	var result string
	if d < 20 {
		result = fmt.Sprintf("%s %s", zeroToNineteen[d], thousand)
	} else if d < 100 {
		result = fmt.Sprintf("%s %s", underHundred(d), thousand)
	} else {
		result = fmt.Sprintf("%s %s", underThousand(d), thousand)
	}
	return result + remainder(" and ", r)
}

func underThousand(n int) string {
	var hundred = "hundred"
	r := n % 100
	d := n / 100
	return fmt.Sprintf("%s %s", zeroToNineteen[d], hundred) + remainder(" and ", r)
}

func underHundred(n int) string {
	var twentyToNinety = []string{"twenty", "thirty", "forty", "fifty", "sixty", "seventy", "eighty", "ninety"}
	r := n % 10
	d := (n / 10) - 2
	return twentyToNinety[d] + remainder("-", r)
}

func remainder(separator string, r int) string {
	if r == 0 {
		return ""
	}
	var result = separator
	if r < 20 {
		result += zeroToNineteen[r]
	} else if r < 100 {
		result += underHundred(r)
	} else {
		result += underThousand(r)
	}
	return result
}
