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
	p1(input)
	p2(input)
}

func p1(input []string) {
	var result int
	for i := range input {
		var numbers []int
		var bigger1, bigger2 int
		values := strings.Split(input[i], "")
		for j := range values {
			num, _ := strconv.Atoi(values[j])
			numbers = append(numbers, num)
		}
		for k := 0; k < len(numbers); k++ {
			if numbers[k] > bigger1 && k != len(numbers)-1 {
				bigger1 = numbers[k]
				bigger2 = numbers[k+1]
			} else if numbers[k] > bigger2 {
				bigger2 = numbers[k]
			}
		}
		result += bigger1*10 + bigger2
	}
	fmt.Println("Result Part 1:", result)
}

func p2(input []string) {
	var result int
	for i := range input {
		var numbers []int
		var positions []int
		values := strings.Split(input[i], "")
		for j := range values {
			num, _ := strconv.Atoi(values[j])
			numbers = append(numbers, num)
		}
		posStart := 0
		posEnd := len(numbers) - 12
		for range 12 {
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
			positions = append(positions, pos)
		}
		var value int
		n := 11
		for _, m := range positions {
			value += numbers[m] * int(math.Pow10(n))
			n--
		}
		result += value
	}
	fmt.Println("Result Part 2:", result)
}
