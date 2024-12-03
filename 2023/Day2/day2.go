//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)
/*
func isLegit(n, R, G, B int, col string) bool{
	if ((col == "red" && n > R) || (col == "blue" && n > B) || col == "green" && n > G){
		return false
	}
	return true
} */

func main(){
	//greenMax := 13
	//redMax := 12
	//blueMax := 14
	sum := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		//count := 0
		//countTot := 0
		riga := scanner.Text()
		red := 0
		blue := 0
		green := 0
		//id, _:= strconv.Atoi(riga[strings.Index(riga, "e")+2:strings.Index(riga, ":")])
		games := riga[strings.Index(riga, ":")+2:]
		//fmt.Println(games)
		arrayGames := strings.Split(games, ";")
		for i:=0; i < len(arrayGames); i++{
			game := strings.Split(arrayGames[i], ",")
			//fmt.Println(game)
			for j:=0; j < len(game); j++{
				values := strings.Split(game[j], " ")
				if values[0] == ""{
					values = values[1:]
				}
				value, _ := strconv.Atoi(values[0]) 
				color := values[1]
				//fmt.Println("color", color, "has value", value)
				if color == "red"{
					if value > red{
						red = value
					}
				}else if color == "blue"{
					if value > blue{
						blue = value
					}
				}else{
					if value > green{
						green = value
					}
				}
				/*
				countTot++
				if isLegit(value, redMax, greenMax, blueMax, color){
					count++
				} */
			}
			//fmt.Println("I valori minimi sono: red", red, "blue", blue, "green", green)
		}
		//fmt.Println("I valori minimi sono: red", red, "blue", blue, "green", green)
		pow := red * blue * green
		//fmt.Println("power: ", pow)
		sum += pow
		/*
		if count == countTot{
			fmt.Println("Partita", id, "accettabile!")
			sum += id
		}else{
			fmt.Println("Partita", id, "non accettabile!")
		}
		count = 0
		countTot = 0
		*/
	}
	fmt.Println("Somma totale: ", sum)
}