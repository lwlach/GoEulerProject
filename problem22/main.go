package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

func main() {
	//var letters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	//for i, r := range letters {
	//	fmt.Println(i+1, alphabeticalPosition(r), string(r))
	//}

	names := readNames()
	fmt.Println(names[5162])
	sort.Strings(names)
	var sumScores int
	for i, name := range names {
		var charSum int
		for _, r := range name {
			charSum += alphabeticalPosition(r)
		}
		score := charSum * (i + 1)
		sumScores += score
	}
	fmt.Println(sumScores)
}

func alphabeticalPosition(r rune) int {
	return int(r - 'A' + 1)
}

func readNames() []string {
	file, err := os.Open("problem22/names.txt")
	if err != nil {
		panic(err)
	}

	reader := bufio.NewReader(file)
	line, err := reader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	line = strings.ReplaceAll(line, "\"", "")
	line = strings.ReplaceAll(line, "\n", "")
	return strings.Split(line, ",")
}
