//Autore: Francesco Corrado
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"unicode"
)

func isSymbol(r rune) bool {
	return r == '*' || r == '#' || r == '+' || r == '$'
}

func sumPartNumbers(schematic []string) int {
	rows := len(schematic)
	sum := 0

	for i := 0; i < rows; i++ {
		for j := 0; j < len(schematic[i]); j++ {
			if unicode.IsDigit(rune(schematic[i][j])) {
				// Check all adjacent cells for symbols
				isPartNumber := false
				for di := -1; di <= 1; di++ {
					for dj := -1; dj <= 1; dj++ {
						if di == 0 && dj == 0 {
							continue
						}
						ni, nj := i+di, j+dj
						if ni >= 0 && ni < rows && nj >= 0 && nj < len(schematic[ni]) {
							if isSymbol(rune(schematic[ni][nj])) {
								isPartNumber = true
								break
							}
						}
					}
					if isPartNumber {
						break
					}
				}
				if isPartNumber {
					// Find the full number
					start := j
					for start > 0 && unicode.IsDigit(rune(schematic[i][start-1])) {
						start--
					}
					end := j
					for end < len(schematic[i])-1 && unicode.IsDigit(rune(schematic[i][end+1])) {
						end++
					}
					if num, err := strconv.Atoi(schematic[i][start : end+1]); err == nil {
						sum += num
					}
					// Skip the rest of the digits in the current number
					j = end
				}
			}
		}
	}
	return sum
}

func readSchematicFromStdin() ([]string, error) {
	var schematic []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		schematic = append(schematic, scanner.Text())
	}
	return schematic, scanner.Err()
}

func main() {
	schematic, err := readSchematicFromStdin()
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	fmt.Println(sumPartNumbers(schematic))
}
