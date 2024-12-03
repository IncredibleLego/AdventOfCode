//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func controlRows(x int, y int, table [][]int, lowPoints []int) []int{
	valx := 0
	if x == 0{
		valx = x+1
	}else{
		valx = x-1
	}

	valy := 0
	if y == 0{
		valy = y+1
	}else{
		valy = y-1
	}

	if y < 0{
		for i:=0; i < len(table[x]); i++{
			a := table[x][i]
			if i == len(table[x])-1{
				if a < table[x][i-1] && a < table[valx][i]{
					lowPoints = append(lowPoints, a)
				}
			}else if i == 0{
				if a < table[x][i+1] && a < table[valx][i]{
					lowPoints = append(lowPoints, a)
				}
			}else{
				if a < table[x][i-1] && a < table[x][i+1] && a < table[valx][i]{
					lowPoints = append(lowPoints, a)
				}
			}
		}
	}else{
		for i:=1; i < len(table)-2; i++{
			a := table[i][y]
			if i == len(table)-1{
				if a < table[i][y-1] && a < table[i][valy]{
					lowPoints = append(lowPoints, a)
				}
			}else{
				if a < table[i-1][y] && a < table[i+1][y] && a < table[i][valy]{
					lowPoints = append(lowPoints, a)
				}
			}
		}
	}
	return lowPoints
}

func main(){
	var sum int
	var table [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), "")
		var temp []int
		for i:=0; i < len(line); i++{
			k, _ := strconv.Atoi(line[i])
			temp = append(temp, k)
		}
		table = append(table, temp)
	}
	fmt.Println("TABELLA:")
	for i:=0; i < len(table); i++{
		fmt.Println(table[i])
	}

	var lowPoints []int
	//Controlla valori centrali
	for x:= 1; x < len(table)-1; x++{
		for y := 1; y < len(table[0])-1; y++{
			a := table[x][y]
			if a < table[x][y+1] && a < table[x][y-1] && a < table[x+1][y] && a < table[x-1][y]{
				lowPoints = append(lowPoints, a)
			}

		}
	}
	//Controlla cornice
	lowPoints = controlRows(0, -1, table, lowPoints)
	lowPoints = controlRows(len(table)-1, -1, table, lowPoints)
	lowPoints = controlRows(-1, 0, table, lowPoints)
	lowPoints = controlRows(-1, len(table[0])-1, table, lowPoints)

	fmt.Println(lowPoints)

	for i:=0; i < len(lowPoints); i++{
		sum += lowPoints[i] + 1
	}
	fmt.Println("SOMMA FINALE:", sum)
}