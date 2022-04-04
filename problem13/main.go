package main

import (
	"bufio"
	"fmt"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	var sum = big.NewInt(0)
	for _, n := range readNumbers() {
		sum.Add(sum, n)
	}
	fmt.Println(sum.String()[:10])
}

func readNumbers() []*big.Int {
	file, err := os.Open("./problem13/numbers.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var numbers []*big.Int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// EOF
			break
		}
		line = strings.Trim(line, "\n")

		// create 50-digit number from string
		var number = new(big.Int)
		for i, nStr := range line {
			n, err := strconv.ParseInt(string(nStr), 10, 64)
			if err != nil {
				panic(err)
			}
			multiplierFactor := int64(len(line) - i)
			multiplier := new(big.Int).Exp(big.NewInt(10), big.NewInt(multiplierFactor), big.NewInt(0))
			number.Add(number, new(big.Int).Mul(big.NewInt(n), multiplier))
		}
		numbers = append(numbers, number)
	}
	return numbers
}
