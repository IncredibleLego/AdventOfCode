//Autore: Francesco Corrado

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main(){
	p1()
}

//Aiutato da Copilot
func findResult(final int, numbers []int) bool {
	return checkCombinations(final, numbers, 0, numbers[0])
}

//Aiutato da Copilot
func checkCombinations(final int, numbers []int, index int, current int) bool {
	if index == len(numbers)-1 {
		return current == final
	}
	nextIndex := index + 1
	return checkCombinations(final, numbers, nextIndex, current+numbers[nextIndex]) || checkCombinations(final, numbers, nextIndex, current*numbers[nextIndex])
}

func p1(){
	totalCalibration := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		fmt.Println(line)
		value, _ := strconv.Atoi(line[:strings.Index(line, ":")])
		//fmt.Println(value)
		num := strings.Split(line[strings.Index(line, ":")+2:], " ")
		var numbers []int
		for i:=0; i<len(num); i++{
			n, _ := strconv.Atoi(num[i])
			numbers = append(numbers, n)
		}
		//fmt.Println(numbers)
		if findResult(value, numbers){
			totalCalibration += value
		}
	}
	fmt.Println("Total calibration:", totalCalibration)
}

/*
1 : 2 operations
2 : 4 operations
3 : 8 operations

2^1
2^2
2^3
*/