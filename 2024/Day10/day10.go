//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func main(){
	var grid [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), "")
		nums := []int{}
		for i := 0; i < len(line); i++{
			x, _ := strconv.Atoi(line[i])
			nums = append(nums, x)
		}
		grid = append(grid, nums)
	}
	p1(grid)
	p2(grid)
}

func checkTrails(grid [][]int, x int, y int, n int, f int) int{
	trails := 0
	if x < 0 || y < 0 || x >= len(grid) || y >= len(grid[x]){
		return 0
	}
	if n == 9 && grid[x][y] == 9{
		if f == 1{
			grid[x][y] = -1
		}
		return 1
	}
	if grid[x][y] == n{
		trails += checkTrails(grid, x+1, y, n+1, f)
		trails += checkTrails(grid, x-1, y, n+1, f)
		trails += checkTrails(grid, x, y+1, n+1, f)
		trails += checkTrails(grid, x, y-1, n+1, f)
	}
	return trails
}

func p1(grid [][]int){
	sumEnding := 0
	for x:=0; x < len(grid); x++{
		for y:=0; y < len(grid[x]); y++{
			if grid[x][y] == 0{
				copia := make([][]int, len(grid))
				for i:=0; i<len(grid); i++{
					copia[i] = make([]int, len(grid[i]))
					copy(copia[i], grid[i])
				}
				_ = checkTrails(copia, x, y, 0, 1)
				for i:=0; i<len(copia); i++{
					for j:=0; j<len(copia[i]); j++{
						if copia[i][j] == -1{
							sumEnding += 1
						}
					}
				}
			}
		}
	}
	fmt.Println("Sum of ending:", sumEnding)	
}

func p2(grid [][]int){
	sumTrails := 0
	for x:=0; x < len(grid); x++{
		for y:=0; y < len(grid[x]); y++{

			if grid[x][y] == 0{
				sumTrails += checkTrails(grid, x, y, 0, 2)
			}
		}
	}
	fmt.Println("Sum of trails:", sumTrails)
}