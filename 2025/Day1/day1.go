//Autore: Francesco Corrado

package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	var input []string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := strings.Split(scanner.Text(), " ")
		input = append(input, text...)
	}
	fmt.Println(input)
	p1(input)
	p2(input)
}

func p1(input []string) {
	var password int
	dial := 50
	for i := range input {
		instruct := input[i][0:1]
		val, _ := strconv.Atoi(input[i][1:])
		val = val % 100
		switch instruct {
		case "R":
			dial += val
		case "L":
			dial -= val
		}
		if dial > 99 {
			dial = dial - 100
		}
		if dial < 0 {
			dial = 100 + dial
		}
		if dial == 0 {
			password++
		}
	}
	fmt.Println("Password Part 1:", password)
}

func p2(input []string) {
	var password int
	dial := 50
	for i := range input {
		instruct := input[i][0:1]
		val, _ := strconv.Atoi(input[i][1:])
		password += val / 100
		val = val % 100
		switch instruct {
		case "R":
			dial += val
		case "L":
			dial -= val
		}
		if dial > 99 {
			if dial != val {
				password++
			}
			dial = dial - 100
		} else if dial < 0 {
			if dial != -val {
				password++
			}
			dial = 100 + dial
		} else if dial == 0 {
			password++
		}
	}
	fmt.Println("Password Part 2:", password)
}
