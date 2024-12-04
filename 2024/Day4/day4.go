//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
)

func main(){
	p1()
	p2()
}

func p1(){
	var table [][]string
	totalXMAS := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		var row []string
		line := scanner.Text()
		for i:=0; i<len(line); i++{
			row = append(row, string(line[i]))
		}
		table = append(table, row)
	}
	for i:=0; i<len(table); i++{
		for j:=0; j<len(table[i]); j++{
			if table[i][j] == "X"{
				if j+3 < len(table[i]) && table[i][j+1] == "M" && table[i][j+2] == "A" && table[i][j+3] == "S"{
					totalXMAS++ //Right check
				}
				if j-3 >= 0 && table[i][j-1] == "M" && table[i][j-2] == "A" && table[i][j-3] == "S"{
					totalXMAS++ //Left check
				}
				if i-3 >=0{
					if table[i-1][j] == "M" && table[i-2][j] == "A" && table[i-3][j] == "S"{
						totalXMAS++ //Up check
					}
					if	j-3 >= 0 && table[i-1][j-1] == "M" && table[i-2][j-2] == "A" && table[i-3][j-3] == "S"{
						totalXMAS++ //Up left check
					}
					if	j+3 < len(table[i]) && table[i-1][j+1] == "M" && table[i-2][j+2] == "A" && table[i-3][j+3] == "S"{
						totalXMAS++ //Up right check
					}
				}
				if i+3 < len(table){
					if table[i+1][j] == "M" && table[i+2][j] == "A" && table[i+3][j] == "S"{
						totalXMAS++ //Down check
					}
					if	j-3 >= 0 && table[i+1][j-1] == "M" && table[i+2][j-2] == "A" && table[i+3][j-3] == "S"{
						totalXMAS++ //Down left check
					}
					if	j+3 < len(table[i]) && table[i+1][j+1] == "M" && table[i+2][j+2] == "A" && table[i+3][j+3] == "S"{
						totalXMAS++ //Down right check
					}
				}
			}
		}
	}
	fmt.Println("XMAS total: ", totalXMAS)
}

func p2(){
	var table [][]string
	totalMAS := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		var row []string
		line := scanner.Text()
		for i:=0; i<len(line); i++{
			row = append(row, string(line[i]))
		}
		table = append(table, row)
	}
	for i:=0; i<len(table); i++{
		for j:=0; j<len(table[i]); j++{
			if table[i][j] == "A"{
				if i-1 >= 0 && i+1 < len(table) && j-1 >= 0 && j+1 < len(table[i]){
					if table[i-1][j-1] == "M" && table[i-1][j+1] == "M" && table[i+1][j-1] == "S" && table[i+1][j+1] == "S"{ //Case 1
						totalMAS++
					}
					if table[i-1][j-1] == "M" && table[i-1][j+1] == "S" && table[i+1][j-1] == "M" && table[i+1][j+1] == "S"{ //Case 2
						totalMAS++
					}
					if table[i-1][j-1] == "S" && table[i-1][j+1] == "S" && table[i+1][j-1] == "M" && table[i+1][j+1] == "M"{ //Case 3
						totalMAS++
					}
					if table[i-1][j-1] == "S" && table[i-1][j+1] == "M" && table[i+1][j-1] == "S" && table[i+1][j+1] == "M"{ //Case 4
						totalMAS++
					}
				}
			}
		}
	}
	fmt.Println("MAS total: ", totalMAS)
}	