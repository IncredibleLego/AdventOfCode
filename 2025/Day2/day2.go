//Autore: Francesco Corrado
//Exec time: 2.4-2.8 sec

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
		text := strings.Split(scanner.Text(), ",")
		input = append(input, text...)
	}
	fmt.Println(input)
	p1(input)
	p2(input)
}

func p1(input []string) {
	var result int
	for i := range input {
		pos := strings.Index(input[i], "-")
		zero, _ := strconv.Atoi(string(input[i][0]))
		if zero == 0 {
			continue
		}
		min, _ := strconv.Atoi(input[i][0:pos])
		max, _ := strconv.Atoi(input[i][pos+1:])
		for j := min; j <= max; j++ {
			number := strconv.Itoa(j)
			if len(number)%2 != 0 {
				continue
			}
			if number[0:len(number)/2] == number[len(number)/2:] {
				result += j
			}
		}
	}
	fmt.Println("Result Part 1:", result)
}

func p2(input []string) {
	var result int
	for i := range input {
		pos := strings.Index(input[i], "-")
		zero, _ := strconv.Atoi(string(input[i][0]))
		if zero == 0 {
			continue
		}
		min, _ := strconv.Atoi(input[i][0:pos])
		max, _ := strconv.Atoi(input[i][pos+1:])
		for j := min; j <= max; j++ {
			number := strconv.Itoa(j)
			div := 1
			for k := 1; k <= len(number)/2; k++ {
				var parti []string
				from := 0
				to := div
				if len(number)%div != 0 {
					div++
					continue
				}
				for l := 0; l < len(number)/div; l++ {
					val := number[from:to]
					parti = append(parti, val)
					from += div
					to += div
				}
				check := true
				for m := 0; m < len(parti)-1; m++ {
					if parti[m] != parti[m+1] {
						check = false
						break
					}
				}
				if check {
					result += j
					break
				}
				div++
			}
		}
	}
	fmt.Println("Result Part 2:", result)
}

/* Optimized version of p2 suggested by Copilot
func p2(input []string) {
	var result int
	for i := range input {
		pos := strings.Index(input[i], "-")
		if input[i][0] == '0' {
			continue
		}
		min, _ := strconv.Atoi(input[i][0:pos])
		max, _ := strconv.Atoi(input[i][pos+1:])
		for j := min; j <= max; j++ {
			s := strconv.Itoa(j)
			n := len(s)
			// un'unica cifra non può essere ripetuta in più parti
			if n < 2 {
				continue
			}
			for d := 1; d <= n/2; d++ {
				if n%d != 0 {
					continue
				}
				if strings.Repeat(s[:d], n/d) == s {
					result += j
					break
				}
			}
		}
	}
	fmt.Println("Result Part 2:", result)
}
*/
