package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	grid := readGrid()
	var highestProduct int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			var product int
			// down
			if i+3 <= len(grid)-1 {
				product = grid[i+3][j] * grid[i+2][j] * grid[i+1][j] * grid[i][j]
				if product > highestProduct {
					highestProduct = product
				}
			}
			// right
			if j+3 <= len(grid[i])-1 {
				product = grid[i][j+3] * grid[i][j+2] * grid[i][j+1] * grid[i][j]
				if product > highestProduct {
					highestProduct = product
				}
			}
			// right diagonal
			if i+3 <= len(grid)-1 && j+3 <= len(grid[i])-1 {
				product = grid[i+3][j+3] * grid[i+2][j+2] * grid[i+1][j+1] * grid[i][j]
				if product > highestProduct {
					highestProduct = product
				}
			}
			// left diagonal
			if i+3 <= len(grid)-1 && j-3 >= 0 {
				product = grid[i+3][j-3] * grid[i+2][j-2] * grid[i+1][j-1] * grid[i][j]
				if product >= highestProduct {
					highestProduct = product
				}
			}
		}
	}
	fmt.Println(highestProduct)
}

func readGrid() [][]int {
	file, err := os.Open("./problem11/grid.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reader := bufio.NewReader(file)
	var grid [][]int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// EOF
			break
		}
		line = strings.Trim(line, "\n")
		numbersStr := strings.Split(line, " ")
		var arr []int
		for _, s := range numbersStr {
			n, _ := strconv.Atoi(s)
			arr = append(arr, n)
		}
		grid = append(grid, arr)
	}
	return grid
}
