//Autore: Francesco Corrado
// Exec time: part1 0.012 sec, part2 0.017 sec

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
	//p1()
	p2()
}

func p1() {
	var result int
	var input [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, strings.Fields(scanner.Text()))
	}
	for k := range input[0] {
		var numbers []int
		var total int
		operation := input[len(input)-1][k]
		for i := 0; i < len(input)-1; i++ {
			val, _ := strconv.Atoi(input[i][k])
			numbers = append(numbers, val)
		}
		if operation == "+" {
			for i := range numbers {
				total += numbers[i]
			}
		} else {
			for i := 0; i < len(numbers); i++ {
				if i == 0 {
					total += numbers[0] * numbers[1]
					i++
				} else {
					total *= numbers[i]
				}
			}
		}
		result += total
	}
	fmt.Println("Result Part 1:", result)
}

func p2() {
	var result int
	var input [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		input = append(input, strings.Split(scanner.Text(), ""))
	}
	for pos := 0; pos < len(input[0]); pos++ {
		var numbers []int
		var operation string
		if input[len(input)-1][pos] != " " {
			operation = input[len(input)-1][pos]
		}
		for {
			var numero []int
			for j := 0; j < len(input)-1; j++ {
				n, err := strconv.Atoi(input[j][pos])
				if err == nil {
					numero = append(numero, n)
				}
			}
			if numero == nil {
				break
			} else {
				var num int
				for l := range numero {
					num += numero[l] * int(math.Pow10(len(numero)-l-1))
				}
				numbers = append(numbers, num)
				pos++
			}
			if pos == len(input[0]) {
				break
			}
		}
		var total int
		if operation == "+" {
			for i := range numbers {
				total += numbers[i]
			}
		} else {
			for i := 0; i < len(numbers); i++ {
				if i == 0 {
					total += numbers[0] * numbers[1]
					i++
				} else {
					total *= numbers[i]
				}
			}
		}
		result += total
	}
	fmt.Println("Result Part 2:", result)
}
