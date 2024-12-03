//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

//PUNTO 2
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan(){
		var progression [][]int
		var dataset []int
		linea := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(linea); i++{
			num, _ := strconv.Atoi(linea[i])
			dataset = append(dataset, num)
		}
		progression = append(progression, dataset)
		fmt.Println(progression)

		//var temporary []int
		for i := 0; i < len(progression); i++{
			var temporary []int
			arr := progression[i]
			for j := 1; j < len(arr); j++{
				temporary = append(temporary, arr[j] - arr[j-1])
			}
			counter0 := 0
			for j := 0; j < len(temporary); j++{
				if temporary[j] == 0{
					counter0++
				}
			}
			progression = append(progression, temporary)
			if counter0 == len(temporary){
				fmt.Println("Terminato")
				break
			}
		}
		for i := 0; i < len(progression); i++{
			fmt.Println(progression[i]) 
		}
		fmt.Println("Elaborazione")
		for i := len(progression)-1; i >= 0; i--{
			arr1 := progression[i] //Array corrente su cui si lavora
			if i == len(progression)-1{
				arr1New := make([]int, 1)
				arr1New[0] = 0
				arr1New = append(arr1New, arr1...)
				progression[i] = arr1New
			}else{
				arr2 := progression[i+1]
				arr1New := make([]int, 1)
				arr1New[0] = arr1[0] - arr2[0]
				arr1New = append(arr1New, arr1...)
				progression[i] = arr1New
			}
		}
		for i := 0; i < len(progression); i++{
			fmt.Println(progression[i]) 
		}
		fmt.Println()
		arr := progression[0]
		sum += arr[0]
	}
	fmt.Println("SOMMA TOTALE: ", sum)
}


/* PUNTO 1
func main(){
	scanner := bufio.NewScanner(os.Stdin)
	sum := 0
	for scanner.Scan(){
		var progression [][]int
		var dataset []int
		linea := strings.Split(scanner.Text(), " ")
		for i := 0; i < len(linea); i++{
			num, _ := strconv.Atoi(linea[i])
			dataset = append(dataset, num)
		}
		progression = append(progression, dataset)
		fmt.Println(progression)

		//var temporary []int
		for i := 0; i < len(progression); i++{
			var temporary []int
			arr := progression[i]
			for j := 1; j < len(arr); j++{
				temporary = append(temporary, arr[j] - arr[j-1])
			}
			counter0 := 0
			for j := 0; j < len(temporary); j++{
				if temporary[j] == 0{
					counter0++
				}
			}
			progression = append(progression, temporary)
			if counter0 == len(temporary){
				fmt.Println("Terminato")
				break
			}
		}
		for i := 0; i < len(progression); i++{
			fmt.Println(progression[i]) 
		}
		fmt.Println("Elaborazione")
		for i := len(progression)-1; i >= 0; i--{
			arr1 := progression[i] //Array corrente su cui si lavora
			if i == len(progression)-1{
				arr1 = append(arr1, 0)
				progression[i] = arr1
			}else{
				arr2 := progression[i+1] //Array precedente
				arr1 = append(arr1, arr1[len(arr1)-1] + arr2[len(arr2)-1])
				progression[i] = arr1
			}
		}
		for i := 0; i < len(progression); i++{
			fmt.Println(progression[i]) 
		}
		fmt.Println()
		arr := progression[0]
		sum += arr[len(arr)-1]
	}
	fmt.Println("SOMMA TOTALE: ", sum)
}
*/