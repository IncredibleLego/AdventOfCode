//Autore: Francesco Corrado
//Exec time: 0.013-15 sec (both parts)

package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strings"
)

func main() {
	var input [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}
	p1(input)
	p2(input)
}

func p1(input [][]string) {
	var result int
	var beams []string
	for range len(input[0]) {
		beams = append(beams, ".")
	}
	beams[slices.Index(input[0], "S")] = "|"
	for i := 2; i < len(input); i += 2 {
		var newBeams []int
		for k := range input[i] {
			if input[i][k] == "^" && beams[k] == "|" {
				result++
				newBeams = append(newBeams, k)
			}
		}
		for _, n := range newBeams {
			beams[n] = "."
			beams[n-1], beams[n+1] = "|", "|"
		}
	}
	fmt.Println("Result Part 1:", result)
}

func p2(input [][]string) {
	var result int
	countBeams := make([]int, len(input[0]))
	countBeams[slices.Index(input[0], "S")] = 1
	for i := 2; i < len(input); i += 2 {
		for k := range input[i] {
			if input[i][k] == "^" {
				countBeams[k-1] += countBeams[k]
				countBeams[k+1] += countBeams[k]
				countBeams[k] = 0
			}
		}
	}
	for l := range countBeams {
		result += countBeams[l]
	}
	fmt.Println("Result Part 2:", result)
}
