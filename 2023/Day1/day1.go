//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"unicode"
	"strconv"
	"strings"
	"sort"
)

func main(){
	var somma, x, y int
	var righe []string

	var mappa map[string]string
	mappa = map[string]string{
		"one": "1",
		"two": "2",
		"three": "3",
		"four": "4",
		"five": "5",
		"six": "6",
		"seven": "7",
		"eight": "8",
		"nine": "9",
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		precedenza := make(map[int]string)
		riga := scanner.Text()
		//fmt.Println("Riga: ", riga)
		for i, _ := range mappa{
			if strings.Contains(riga, i){
				precedenza[strings.Index(riga, i)] = i
			}
		}
		keys := make([]int, 0, len(precedenza))
		//fmt.Println(precedenza)

		for k := range precedenza{
			keys = append(keys, k)
		}
		sort.Ints(keys)
		//fmt.Println("KEYS: ", keys)

		if len(keys) > 0{
			out := riga[:keys[0]]
			for i:=0; i<len(keys); i++{
				val := precedenza[keys[i]]
				if len(out) < keys[i]{
					out += riga[keys[i-1]:keys[i]]
				}
				out += mappa[val]
			}
			if len(out) < len(riga){
				out += riga[keys[len(keys)-1]:]
			}
			//fmt.Println("Out: ", out)
			righe = append(righe, out)
		}else{
			righe = append(righe, riga)
		}
	}

	for i:=0; i<len(righe); i++{

		riga := righe[i]
		for i:=0; i<len(riga); i++{
			if unicode.IsDigit(rune(riga[i])){
				x, _ = strconv.Atoi(string(riga[i]))
				break
			}
		}
		for i:=len(riga)-1; i>=0; i--{
			if unicode.IsDigit(rune(riga[i])){
				y, _ = strconv.Atoi(string(riga[i]))
				break
			}
		}
		fmt.Println("Riga", i+1, ":", riga, "Valore:", x*10 + y)
		somma += x*10 + y
	}
	fmt.Println("Somma: ", somma)
}

/*
for _, v := range keys {
			val := precedenza[v]
			fmt.Println("Valore: ", val)
			fmt.Println("Riga prima: ", riga)
			riga = strings.ReplaceAll(riga, val, mappa[val])
			fmt.Println("Riga dopo: ", riga)
		}
		righe = append(righe, riga)
	*/