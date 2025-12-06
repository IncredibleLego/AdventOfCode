//Autore: Francesco Corrado
//Exec time: 0.06 sec

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var fresh []string
	var ids []int
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		text := scanner.Text()
		fresh = append(fresh, text)
	}
	for scanner.Scan() {
		id, _ := strconv.Atoi(scanner.Text())
		ids = append(ids, id)

	}
	p1(fresh, ids)
	p2(fresh)
}

func p1(fresh []string, ids []int) {
	var result int
	for k := range ids {
		for i := range fresh {
			values := strings.Split(fresh[i], "-")
			min, _ := strconv.Atoi(values[0])
			max, _ := strconv.Atoi(values[1])
			if ids[k] >= min && ids[k] <= max {
				result++
				break
			}
		}
	}
	fmt.Println("Result Part 1:", result)
}

func p2(fresh []string) {
	var result int
	var intros []int
	for i := range fresh {
		values := strings.Split(fresh[i], "-")
		min, _ := strconv.Atoi(values[0])
		intros = append(intros, min)
	}
	for {
		var changed bool
		fresh, changed = reduceIntervals(fresh)
		if changed != true {
			break
		}
	}
	for i := range fresh {
		intervals := strings.Split(fresh[i], "-")
		min, _ := strconv.Atoi(intervals[0])
		max, _ := strconv.Atoi(intervals[1])
		result += max + 1 - min
	}
	fmt.Println("Result Part 2:", result)
}

func reduceIntervals(fresh []string) ([]string, bool) {
	sort.Slice(fresh, func(i, j int) bool {
		ai, _ := strconv.Atoi(strings.Split(fresh[i], "-")[0])
		aj, _ := strconv.Atoi(strings.Split(fresh[j], "-")[0])
		return ai < aj
	})
	var reduced []string
	changed := false
	i := 0
	for i < len(fresh)-1 {
		interval1 := strings.Split(fresh[i], "-")
		interval2 := strings.Split(fresh[i+1], "-")
		min1, _ := strconv.Atoi(interval1[0])
		max1, _ := strconv.Atoi(interval1[1])
		min2, _ := strconv.Atoi(interval2[0])
		max2, _ := strconv.Atoi(interval2[1])
		if min2 <= max1+1 {
			newMin := min1
			newMax := max1
			if max2 > max1 {
				newMax = max2
			}
			merged := fmt.Sprintf("%d-%d", newMin, newMax)
			reduced = append(reduced, merged)
			changed = true
			i += 2
			continue
		}
		reduced = append(reduced, fresh[i])
		i++
	}
	if i == len(fresh)-1 {
		reduced = append(reduced, fresh[i])
	}
	return reduced, changed
}

//Risultato: 342433357244012

//309290793217139
