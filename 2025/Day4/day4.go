//Autore: Francesco Corrado
//Exec time: 0.130 sec

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	var input [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), "")
		input = append(input, text)
	}
	p1(input)
	p2(input)
}

func p1(input [][]string) {
	var result int
	for x := range input {
		for y := range input[x] {
			if input[x][y] == "@" {
				countRoll := checkNearby(input, x, y)
				if countRoll < 4 {
					result++
				}
			}
		}
	}
	fmt.Println("Result Part 1:", result)
}

func p2(input [][]string) {
	var result int
	for {
		var cpy [][]string
		var control int
		for x := range input {
			var tempArray []string
			for y := range input[x] {
				tempArray = append(tempArray, input[x][y])
				if input[x][y] == "@" {
					countRoll := checkNearby(input, x, y)
					if countRoll < 4 {
						tempArray[y] = "x"
						result++
						control++
					}
				}
			}
			cpy = append(cpy, tempArray)
		}
		input = cpy
		if control == 0 {
			break
		}
	}
	fmt.Println("Result Part 2:", result)
}

func checkNearby(input [][]string, x, y int) int {
	var countRoll int
	if input[x][y] == "@" {
		if x-1 >= 0 {
			if input[x-1][y] == "@" {
				countRoll++
			}
			if y-1 >= 0 {
				if input[x-1][y-1] == "@" {
					countRoll++
				}
			}
			if y+1 < len(input[x]) {
				if input[x-1][y+1] == "@" {
					countRoll++
				}
			}
		}
		if x+1 < len(input) {
			if input[x+1][y] == "@" {
				countRoll++
			}
			if y-1 >= 0 {
				if input[x+1][y-1] == "@" {
					countRoll++
				}
			}
			if y+1 < len(input[x]) {
				if input[x+1][y+1] == "@" {
					countRoll++
				}
			}
		}
		if y-1 >= 0 {
			if input[x][y-1] == "@" {
				countRoll++
			}
		}
		if y+1 < len(input[x]) {
			if input[x][y+1] == "@" {
				countRoll++
			}
		}
	}
	return countRoll
}
