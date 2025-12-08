//Autore: Francesco Corrado
//Exec time: part1 0,98 sec, part2 2.2 sec, total 3 sec

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var input [][]int
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		var num []int
		point := strings.Split(scanner.Text(), ",")
		for r := range 3 {
			n, _ := strconv.Atoi(point[r])
			num = append(num, n)
		}
		input = append(input, num)
	}
	p1(input)
	p2(input)
}

func p1(input [][]int) {
	var result int
	distances := make(map[float64][][]int)
	for i := range input {
		for k := i + 1; k < len(input); k++ {
			distance := math.Sqrt(math.Pow(float64(input[i][0]-input[k][0]), 2) + math.Pow(float64(input[i][1]-input[k][1]), 2) + math.Pow(float64(input[i][2]-input[k][2]), 2))
			distances[distance] = [][]int{input[i], input[k]}
		}
	}
	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	sort.Float64s(keys)
	var connections [][][]int

	x := 0
	for _, k := range keys {
		check1 := isInConnection(distances[k][0], connections)
		check2 := isInConnection(distances[k][1], connections)

		if check1 == -1 && check2 == -1 {
			connections = append(connections, [][]int{distances[k][0], distances[k][1]})
		} else if check1 > -1 && check2 > -1 && check1 != check2 {
			connections[check1] = append(connections[check1], connections[check2]...)
			connections = append(connections[:check2], connections[check2+1:]...)
		} else if check1 > -1 && check2 == -1 {
			connections[check1] = append(connections[check1], distances[k][1])
		} else if check1 == -1 && check2 > -1 {
			connections[check2] = append(connections[check2], distances[k][0])
		}
		if x == 999 {
			break
		}
		x++
	}

	var lenghts []int
	for i := range connections {
		lenghts = append(lenghts, len(connections[i]))
	}
	sort.Sort(sort.Reverse(sort.IntSlice(lenghts)))
	result = lenghts[0] * lenghts[1] * lenghts[2]
	fmt.Println("Result Part 1:", result)
}

func p2(input [][]int) {
	var result int

	distances := make(map[float64][][]int)
	for i := range input {
		for k := i + 1; k < len(input); k++ {
			distance := math.Sqrt(math.Pow(float64(input[i][0]-input[k][0]), 2) + math.Pow(float64(input[i][1]-input[k][1]), 2) + math.Pow(float64(input[i][2]-input[k][2]), 2))
			distances[distance] = [][]int{input[i], input[k]}
		}
	}
	keys := make([]float64, 0, len(distances))
	for k := range distances {
		keys = append(keys, k)
	}
	sort.Float64s(keys)

	var connections [][][]int
	var previousConnections int
	var placeLastConnection [][]int

	tempConnections := make(map[string]bool)
	for _, m := range keys {
		val1 := distances[m][0]
		val2 := distances[m][1]
		key1 := fmt.Sprintf("%d,%d,%d", val1[0], val1[1], val1[2])
		key2 := fmt.Sprintf("%d,%d,%d", val2[0], val2[1], val2[2])
		tempConnections[key1] = true
		tempConnections[key2] = true
	}

	for r := range tempConnections {
		parts := strings.Split(r, ",")
		coord := make([]int, 3)
		for i := range 3 {
			coord[i], _ = strconv.Atoi(parts[i])
		}
		connections = append(connections, [][]int{coord})
	}

	for _, k := range keys {
		check1 := isInConnection(distances[k][0], connections)
		check2 := isInConnection(distances[k][1], connections)
		if check1 == -1 && check2 == -1 {
			connections = append(connections, [][]int{distances[k][0], distances[k][1]})
		} else if check1 > -1 && check2 > -1 && check1 != check2 {
			connections[check1] = append(connections[check1], connections[check2]...)
			connections = append(connections[:check2], connections[check2+1:]...)
		} else if check1 > -1 && check2 == -1 {
			connections[check1] = append(connections[check1], distances[k][1])
		} else if check1 == -1 && check2 > -1 {
			connections[check2] = append(connections[check2], distances[k][0])
		}

		if previousConnections > 1 && len(connections) == 1 {
			placeLastConnection = distances[k]
			break
		}
		previousConnections = len(connections)
	}

	result = placeLastConnection[0][0] * placeLastConnection[1][0]
	fmt.Println("Result Part 2:", result)
}

func isInConnection(coordinates []int, connections [][][]int) int {
	for i := range connections {
		for k := range connections[i] {
			if connections[i][k][0] == coordinates[0] && connections[i][k][1] == coordinates[1] && connections[i][k][2] == coordinates[2] {
				return i
			}
		}
	}
	return -1
}
