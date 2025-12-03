//Autore: Francesco Corrado

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		input = append(input, text)
	}
	day3(input, 2, 1)
	day3(input, 12, 2)
}

func day3(input []string, x int, part int) {
	var result int
	for i := range input {
		var numbers []int
		values := strings.Split(input[i], "")
		for j := range values {
			num, _ := strconv.Atoi(values[j])
			numbers = append(numbers, num)
		}
		posStart := 0
		posEnd := len(numbers) - x
		n := x - 1
		for range x {
			bigger := 0
			pos := 0
			for l := posStart; l <= posEnd; l++ {
				if numbers[l] > bigger {
					bigger = numbers[l]
					pos = l
				}
				if bigger == 9 {
					break
				}
			}
			posStart = pos + 1
			posEnd++
			result += bigger * int(math.Pow10(n))
			n--
		}
	}
	fmt.Println("Result Part", part, ":", result)
}
