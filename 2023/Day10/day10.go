//Autore: Francesco Corrado

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
)

func path(table [][]string, x int, y int, count int, from string) int{
	if table[x][y] == "."{
		fmt.Println("Posizione finale:", table[x][y], "a coordinate:", x, y)
		return count
	}else{
		fmt.Println("Posizione attuale:", table[x][y], "a coordinate:", x, y, "from:", from)
		switch table[x][y]{
		case "|":
			if from == "n"{
				fmt.Println("Next pos: (", table[x+1][y], x+1, y, ")")
				return path(table, x+1, y, count+1, "n")
				//return 0
			}else{
				fmt.Println("Next pos: (", table[x-1][y], x-1, y , ")")
				return path(table, x-1, y, count+1, "s")
				//return 5
			}
		case "-":
			if from == "w"{
				fmt.Println("Next pos: (", table[x][y+1], x, y+1 , ")")
				return path(table, x, y+1, count+1, "w")
			}else{
				fmt.Println("Next pos: (", table[x][y-1] , x, y-1 , ")")
				return path(table, x, y-1, count+1, "e")
			}
		case "L":
			if from == "n"{
				//fmt.Println("PROVA")
				fmt.Println("Next pos: (", table[x][y+1], x, y+1 , ")")
				return path(table, x, y+1, count+1, "w")
			}else{
				fmt.Println("Next pos: (", table[x-1][y], x-1, y , ")")
				return path(table, x-1, y, count+1, "s")
			}
		case "J":
			if from == "n"{
				fmt.Println("Next pos: (", table[x][y-1] , x, y-1 , ")")
				return path(table, x, y-1, count+1, "e")
			}else{
				fmt.Println("Next pos: (", table[x-1][y] , x-1, y , ")")
				return path(table, x-1, y, count+1, "s")
			}
		case "7":
			if from == "s"{
				fmt.Println("Next pos: (", table[x][y-1] , x, y-1 , ")")
				return path(table, x, y-1, count+1, "e")
			}else{
				fmt.Println("Next pos: (", table[x+1][y] , x+1, y , ")")
				return path(table, x+1, y, count+1, "n")
			}
		case "F":
			if from == "s"{
				fmt.Println("Next pos: (", table[x][y+1] , x, y+1 , ")")
				return path(table, x, y+1, count+1, "w")
			}else{
				fmt.Println("Next pos: (", table[x+1][y] , x+1, y , ")")
				return path(table, x+1, y, count+1, "n")
			}
		default:
			return count
		}
	}
}

func main(){
	var table [][]string
	scanner := bufio.NewScanner(os.Stdin)
	k := 0
	x := 0
	y := 0
	for scanner.Scan(){
		row := scanner.Text()
		pos := strings.Index(row, "S")
		if pos != -1{
			y = pos
			x = k
		}
		dividedRow := strings.Split(row, "")
		table = append(table, dividedRow)
		k++
	}
	fmt.Println("Posizione iniziale:", table[x][y], "a coordinate:", x, y)

	for i:=0; i < len(table); i++{
		fmt.Println(table[i])
	}

	dist := 0
	if x-1 >= 0{
		fmt.Println("\nNord:", table[x-1][y])
		tempDist := path(table, x-1, y, 1, "s")
		if tempDist > dist{
			dist = tempDist
		}
	}
	if x+1 < len(table){
		fmt.Println("\nSud:", table[x+1][y])
		tempDist := path(table, x+1, y, 1, "n")
		if tempDist > dist{
			dist = tempDist
		}
	}
	if y-1 >= 0{
		fmt.Println("\nOvest:", table[x][y-1])
		tempDist := path(table, x, y-1, 1, "e")
		if tempDist > dist{
			dist = tempDist
		}
	}
	if y+1 < len(table[0]){
		fmt.Println("\nEst:", table[x][y+1])
		tempDist := path(table, x, y+1, 1, "w")
		if tempDist > dist{
			dist = tempDist
		}
	}
	fmt.Println("Distanza massima:", dist)
	fmt.Println("(diviso 2):", dist/2)
}