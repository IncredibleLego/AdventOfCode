//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"regexp"
	"strconv"
)

func main(){
	totalSum := 0
	final := ""
	//var indexes []int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := scanner.Text()
		final += line
	}
	fmt.Println(final)

	re := regexp.MustCompile(`mul\(\d{1,3},\d{1,3}\)`)
	matches := re.FindAllStringIndex(final, -1)
	for _, match := range matches {
		start, end := match[0], match[1]
		substr := final[start:end]
		//fmt.Println("Trovata stringa:", substr, "in pos", start)

		nums := regexp.MustCompile(`\d+`).FindAllString(substr, -1)
		if len(nums) == 2 {
			num1, _ := strconv.Atoi(nums[0])
			num2, _ := strconv.Atoi(nums[1])
			result := num1 * num2
			totalSum += result
			//fmt.Println("Risultato moltiplicazione:", result)
		}
	}

	fmt.Println("Somma Totale:", totalSum)
}