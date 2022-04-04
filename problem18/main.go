package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	tri := readTriangle()
	for i := len(tri) - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			if tri[i+1][j] > tri[i+1][j+1] {
				tri[i][j] += tri[i+1][j]
				continue
			}
			tri[i][j] += tri[i+1][j+1]
		}
	}
	fmt.Println(tri[0][0])
}

func readTriangle() [][]int {
	file, err := os.Open("problem18/triangle.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	var triangle [][]int
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			// EOF
			break
		}
		line = strings.Trim(line, "\n")
		numbers := strings.Split(line, " ")
		var triangleLine []int
		for _, number := range numbers {
			n, _ := strconv.Atoi(number)
			triangleLine = append(triangleLine, n)
		}
		triangle = append(triangle, triangleLine)
	}
	return triangle
}
