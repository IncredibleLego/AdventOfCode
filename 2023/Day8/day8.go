//Autore: Francesco Corrado

package main

import(
	"fmt"
	"bufio"
	"os"
	"strings"
)

type Nodo struct{
	L, R string
}

func main(){
	//count := 0
	//i := 0
	nodi := make(map[string]Nodo)
	var startNodes []string
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	val := strings.Split(scanner.Text(), "")
	fmt.Println("Valori: ", val)
	scanner.Scan()
	for scanner.Scan(){
		var n Nodo
		line := scanner.Text()
		//fmt.Println(line)
		nome := line[:strings.Index(line, " ")]
		if nome[2] == 'A'{
			startNodes = append(startNodes, nome)
		}
		n.L = line[strings.Index(line, "(")+1:strings.Index(line, ",")]
		n.R = line[strings.Index(line, ",")+2:strings.Index(line, ")")]
		//fmt.Println(n)
		nodi[nome] = n
	}
	//fmt.Println(nodi)
	fmt.Println("startNodes", startNodes)


	k := 0
	results := []int{}
	for j := 0; j < len(startNodes); j++{
		steps := 0
		//c := startNodes[j]
		for{
			c := startNodes[j]
			fmt.Println(startNodes[j])
			fmt.Println(val[k])
			if (c[2] == 'Z'){
				break
			}
			if val[k] == "L"{
				startNodes[j] = nodi[startNodes[j]].L
				fmt.Println(startNodes[j], "è appena diventato", nodi[startNodes[j]].L )
			}else{
				startNodes[j] = nodi[startNodes[j]].R
				fmt.Println(startNodes[j], "è appena diventato", nodi[startNodes[j]].R )
			}

			steps++
			//c = startNodes[j]
			fmt.Println("c", c)
			fmt.Println("scelte: ", nodi[startNodes[j]].L, nodi[startNodes[j]].R)
			if (c[2] == 'Z'){
				break
			}
			//c = startNodes[k]
			k++
			if k == len(val){
				k = 0
			}
		}
		fmt.Println("appeso ", steps)
		results = append(results, steps)
		steps = 0
	}
	out := results[0]
	for i := 1; i < len(results); i++ {
		out = lcm(out, results[i])
	}
	fmt.Println("SOMMA FINALE:", out)



	/*
	name := ""
	for{
		//fmt.Println("PASSAGGIO ", count, "\n")
		//fmt.Println("val[i]:", val[i])
		for j := 0; j < len(startNodes); j++{
			name = startNodes[j]
			//fmt.Println("Nome:", name)
			if val[i] == "L"{
				startNodes[j] = nodi[name].L
			}else{
				startNodes[j] = nodi[name].R
			}
		}
		count++
		fmt.Println("startNodes:", startNodes)
		countZ := 0
		for j := 0; j < len(startNodes); j++{
			n := startNodes[j]
			if n[2] == 'Z'{
				countZ++
			}
		}
		if countZ == len(startNodes){
			break
		}

		i++
		if i == len(val){
			i = 0
		}
		fmt.Println("Passi necessari:", count)

	}
	fmt.Println("Passi necessari:", count) */
}

func gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int) int {
	return a / gcd(a, b) * b
}