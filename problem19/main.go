package main

import (
	"fmt"
	"time"
)

func main() {
	var sundays int
	d := time.Date(1901, 1, 1, 0, 0, 0, 0, time.Local)
	for {
		if d.Weekday() == time.Sunday && d.Day() == 1 {
			sundays++
			d = d.AddDate(0, 0, 28)
			continue
		}
		d = d.AddDate(0, 0, 1)
		if d.Year() == 2000 && d.Month() == 12 && d.Day() == 31 {
			break
		}
	}
	fmt.Println(sundays)
}
