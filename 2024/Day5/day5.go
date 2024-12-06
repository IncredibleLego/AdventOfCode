//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"slices"
)

func checkPage(rules map[int][]int, page []int) bool{
	for i:=1; i<len(page); i++{
		rul := rules[page[i]]
		for j:= i-1; j>=0; j--{
			if slices.Contains(rul, page[j]){
				return false
			}
		}
	}
	return true
}

func orderPages(page []int, rules map[int][]int) []int{
	for{
		check := true
		for i:=1; i<len(page); i++{
			rul := rules[page[i]]
			for j:= i-1; j>=0; j--{
				if slices.Contains(rul, page[j]){
					page[i], page[j] = page[j], page[i]
					check = false
				}
			}
		}
		if check{
			break
		}
	}
	return page
}

//Both parts 1 and 2
func main(){
	var pages [][]int
	totalValid := 0
	totalInvalid := 0
	rules := make(map[int][]int)
	scanner := bufio.NewScanner(os.Stdin)
	check := false
	for scanner.Scan(){
		line := scanner.Text()
		if line == ""{
			check = true
			continue
		}
		if !check{
			n := strings.Split(line, "|")
			k, _ := strconv.Atoi(n[0])
			v, _ := strconv.Atoi(n[1])
			array := rules[k]
			array = append(array, v)
			rules[k] = array
		}else{
			var page []int
			pag := strings.Split(line, ",")
			for i:=0; i<len(pag); i++{
				x, _ := strconv.Atoi(pag[i])
				page = append(page, x)
			}
			pages = append(pages, page)
		}
	}
	for i:=0; i<len(pages); i++{
		if checkPage(rules, pages[i]){
			totalValid += pages[i][len(pages[i])/2]
		}else{
			pages[i] = orderPages(pages[i], rules)
			totalInvalid += pages[i][len(pages[i])/2]
		}
	}
	fmt.Println("Total valid: ", totalValid)
	fmt.Println("Total invalid: ", totalInvalid)
}

