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

//Per eseguire le diverse parti, decommentare le righe segnate come Parte 1 e commentare quelle segnate come Parte 2 o viceversa

func main(){
	var listOne, listTwo []int //Parte 1
	//var listOne, listTwo []string //Parte 2
	//var mappa = make(map[string]int) //Parte 2
	//similarityScore := 0 //Parte 2
	sum := 0 //Parte 1
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := strings.Split(scanner.Text(), " ")
		//listOne = append(listOne, line[0]) //Parte 2
		//listTwo = append(listTwo, line[3]) //Parte 2
		
		x, _ := strconv.Atoi(line[0])
		listOne = append(listOne, x)
		x, _ = strconv.Atoi(line[3])
		listTwo = append(listTwo, x)
	}
	sort.Ints(listOne) //Parte 1
	sort.Ints(listTwo) //Parte 1
	//sort.Strings(listOne) //Parte 2
	//sort.Strings(listTwo) //Parte 2
	/* Parte 2
	for i:=0; i<len(listTwo); i++{
		mappa[listTwo[i]]++
	} */
	
	for i:=0; i<len(listOne); i++{
		if listOne[i] > listTwo[i]{
			sum += listOne[i] - listTwo[i]
			//fmt.Println(listOne[i] - listTwo[i])
		}else{
			sum += listTwo[i] - listOne[i]
			//fmt.Println(listTwo[i] - listOne[i])
		}
	}
	fmt.Println("Somma totale: ", sum) 

	/* Parte 2
	for i:= 0; i<len(listOne); i++{
		x := mappa[listOne[i]]
		n, _ := strconv.Atoi(listOne[i])
		similarityScore += x*n
	}
	fmt.Println("similarityScore: ", similarityScore) */
}