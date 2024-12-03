//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
	"sort"
)

func lineValue(w, h []int, n int) int{
	output := 0
	for i:=0; i < n; i++{
		if i == 0{
			output++
		}else{
			output *= 2
		}
	}
	return output
}

func totalRows(v, c, cou map[int]int, keys []int) int{
	output := 0

	for i:=0; i < len(keys)+1; i++{
		cop := cou[i]

		for k:=0; k < c[i]; k++{

			for j:=i+1; j < (i+1+cop); j++{
				c[j]++
			}
		}
		//fmt.Println("valore riga", i, ":", c[i], "*", v[i], "=", c[i]*v[i])
		output += c[i] * v[i]
	}
	return output
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	values := make(map[int]int)
	copies := make(map[int]int)
	counter := make(map[int]int)
	for scanner.Scan(){
		var winners []int
		var myNumbers []int
		count := 0
		line := scanner.Text()
		g := strings.Replace(line[strings.Index(line, "d")+2:strings.Index(line, ":")], " ", "", -1)
		id, _ := strconv.Atoi(g)
		//fmt.Println(id)
		one := line[strings.Index(line, ":")+2:strings.Index(line, "|")-1]
		two := line[strings.Index(line, "|")+2:]
		win := strings.Split(one, " ")
		for i:=0; i < len(win); i++{
			x, err := strconv.Atoi(win[i])
			if err == nil{
				winners = append(winners, x)
			}
		}
		have := strings.Split(two, " ")
		for i:=0; i < len(have); i++{
			x, err := strconv.Atoi(have[i])
			if err == nil{
				myNumbers= append(myNumbers, x)
			}
		}
		for i:=0; i < len(myNumbers); i++{
			for j :=0; j < len(winners); j++{
				if myNumbers[i] == winners[j]{
					count++
				}
			}
		}
		values[id] = lineValue(winners, myNumbers, count)
		copies[id]++
		counter[id] = count
	}



    keys := make([]int, 0, len(values))
    for k := range values{
        keys = append(keys, k)
    }
    sort.Ints(keys)

	final := totalRows(values, copies, counter, keys)


	/*
	for i:=0; i < len(keys)+1; i++{
		fmt.Println("la riga", i, "vale", values[i])
	} */

	risultato := 0

	for i:=0; i < len(keys)+1; i++{
		//fmt.Println("la riga", i, "ha", copies[i], "copie")
		risultato += copies[i]
	}
	fmt.Println("RISULTATO: ", risultato)


	/*for i:=0; i < len(keys)+1; i++{
		fmt.Println(i, "ha count", counter[i])
	} */

	fmt.Println("SOMMA FINALE: ", final)
}