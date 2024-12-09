//Autore: Francesco Corrado

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"slices"
)

func main() {
	f1()
	//f2()
}

func f1(){
	x, y, i := 0, 0, 0
	direction := "up"
	steps := 0
	var area [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), "")
		if slices.Index(line, "^") != -1{
			x = i
			y = slices.Index(line, "^")
		}
		area = append(area, line)
		i++
	}
	
	for i:=0; i < len(area); i++{
		fmt.Println(area[i])
	}
	fmt.Println("Pos inizio - x:", x, "y:", y)
	area[x][y] = "X"

	check := true
	for check{
		switch direction{
			case "up":
				if x - 1 < 0{
					check = false
					break
				}
				if area[x-1][y] == "#"{
					direction = "right"
				}else{
					x--
				}
			case "right":
				if y + 1 >= len(area[0]){
					check = false
					break
				}
				if area[x][y+1] == "#"{
					direction = "down"
				}else{
					y++
				}
			case "down":
				if x + 1 >= len(area){
					check = false
					break
				}
				if area[x+1][y] == "#"{
					direction = "left"
				}else{
					x++
				}
			case "left":
				if y - 1 < 0{
					check = false
					break
				}
				if area[x][y-1] == "#"{
					direction = "up"
				}else{
					y--
				}
		}
		area[x][y] = "X"
		//fmt.Println("Posizione - x:", x, "y:", y, "direction: ", direction,"steps:", steps)
	}
	for i:=0; i < len(area); i++{
		fmt.Println(area[i])
		for j:=0; j < len(area[i]); j++{
			if area[i][j] == "X"{
				steps++
			}
		}
	}

	fmt.Println("Steps:", steps)
}

func checkLoop(area [][]string, x int, y int, direction string) int{
	switch direction{
		case "up":
			area[x-1][y] = "#"
		case "right":
			area[x][y+1] = "#"
		case "down":
			area[x+1][y] = "#"
		case "left":
			area[x][y-1] = "#"
	}


	i := 0
	check := true
	for check{
		switch direction{
			case "up":
				if x - 1 < 0{
					check = false
					break
				}
				if area[x-1][y] == "#"{
					direction = "right"
				}else{
					x--
				}
			case "right":
				if y + 1 >= len(area[0]){
					check = false
					break
				}
				if area[x][y+1] == "#"{
					direction = "down"
				}else{
					y++
				}
			case "down":
				if x + 1 >= len(area){
					check = false
					break
				}
				if area[x+1][y] == "#"{
					direction = "left"
				}else{
					x++
				}
			case "left":
				if y - 1 < 0{
					check = false
					break
				}
				if area[x][y-1] == "#"{
					direction = "up"
				}else{
					y--
				}
		}
		area[x][y] = "X"
		i++
		if i > 50{
			return 1
		}
	}
	return 0

}

func f2(){
	x, y, i := 0, 0, 0
	direction := "up"
	steps := 0
	totalLoops := 0
	var area [][]string
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), "")
		if slices.Index(line, "^") != -1{
			x = i
			y = slices.Index(line, "^")
		}
		area = append(area, line)
		i++
	}
	
	for i:=0; i < len(area); i++{
		fmt.Println(area[i])
	}
	fmt.Println("Pos inizio - x:", x, "y:", y)
	area[x][y] = "X"

	check := true
	for check{
		switch direction{
			case "up":
				if x - 1 < 0{
					check = false
					break
				}
				if area[x-1][y] == "#"{
					direction = "right"
				}else{
					totalLoops += checkLoop(area, x, y, direction)
					x--
				}
			case "right":
				if y + 1 >= len(area[0]){
					check = false
					break
				}
				if area[x][y+1] == "#"{
					direction = "down"
				}else{
					totalLoops += checkLoop(area, x, y, direction)
					y++
				}
			case "down":
				if x + 1 >= len(area){
					check = false
					break
				}
				if area[x+1][y] == "#"{
					direction = "left"
				}else{
					totalLoops += checkLoop(area, x, y, direction)
					x++
				}
			case "left":
				if y - 1 < 0{
					check = false
					break
				}
				if area[x][y-1] == "#"{
					direction = "up"
				}else{
					totalLoops += checkLoop(area, x, y, direction)
					y--
				}
		}
		area[x][y] = "X"
		//fmt.Println("Posizione - x:", x, "y:", y, "direction: ", direction,"steps:", steps)
	}
	for i:=0; i < len(area); i++{
		fmt.Println(area[i])
		for j:=0; j < len(area[i]); j++{
			if area[i][j] == "X"{
				steps++
			}
		}
	}

	fmt.Println("Steps:", steps)
	fmt.Println("Total Loops:", totalLoops)
}