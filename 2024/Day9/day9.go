//Autore: Francesco Corrado

package main

import(
	"fmt"
	"strconv"
	"strings"
)

func main(){
	var input string
	fmt.Scan(&input)
	fmt.Println("input iniziale: ", input)
	var values []int
	for i:=0; i < len(input); i++{
		x, _ := strconv.Atoi(string(input[i]))
		values = append(values, x)
	}
	p1(values)
	p2(values)
}

func findIndex(s int, c string, values []string, dir bool) int{
	x := s
	for{
		if strings.ContainsAny(values[x], c){
			return x
		}
		if dir{
			x++
		}else{
			x--
		}
	}
}

func p1(values []int){
	var final []string
	var checksum int
	id := 0
	for i:=0; i < len(values); i++{
		if values[i] != 0{
			if i % 2 == 0{
				for j:=0; j < values[i]; j++{
					final = append(final, strconv.Itoa(id))
				}
				id++
			}else{
				for j:=0; j < values[i]; j++{
					final = append(final, ".")
				}
			}
		}
	}
	for{
		a := findIndex(0, ".", final, true)
		b := findIndex(len(final) - 1, "0123456789", final, false)
		if a > b{
			break
		}
		final[a], final[b] = final[b], final[a]
	}
	for i:=0; i < len(final); i++{
		n, err := strconv.Atoi(final[i])
		if err != nil{
			break
		}
		checksum += i * n
	}
	fmt.Println("checksum: ", checksum)
}

func p2(values []int){
	var final []string
	id := 0
	for i:=0; i < len(values); i++{
		if values[i] != 0{
			if i % 2 == 0{
				final = append(final, strings.Repeat(strconv.Itoa(id), values[i]))
				id++
			}else{
				final = append(final, strings.Repeat(".", values[i]))
			}
		}
	}
	fmt.Println("final: ", final)

	//for{
		a := 0
		b := len(final) - 1
		mod := 0
		for{
			a := findIndex(a, ".", final, true)
			a = a + mod
			b := findIndex(b, "0123456789", final, false)
			if a > b{
				break
			}
			fmt.Println("a: ", a, "b: ", b, final[a], final[b])
			if len(final[a]) > len(final[b]){
				l := len(final[a]) - len(final[b])
				temp := make([]string, len(final))
				for i:=0; i < len(final); i++{
					temp[i] = final[i]
				}

				fmt.Println("FINAL INIZIALE: ", final)

				temp[a] = strings.Repeat(".", len(final[b]))
				temp[a], temp[b] = temp[b], temp[a]
				fmt.Println("Changed temp[a]: ", temp[a], "temp[b]: ", temp[b])
				fmt.Println("temp: ", temp, "sub: ", len(final[a]) - len(final[b]))


				temp = append(temp[:a+1], strings.Repeat(".", l))
				temp = append(temp, final[a+1:]...)
				fmt.Println("temp: ", temp, "\nfinal: ", final)
				temp[b+1] = final[b]
				fmt.Println(temp[b], temp[b+1])
				fmt.Println("a: ", a, "b: ", b)


				//fmt.Println("temp: ", temp, "final: ", final)
				final = temp
				fmt.Println("FINAL FINALE  : ", final)
			}else if len(final[a]) < len(final[b]){
				mod = a
				a++
			}else{
				final[a], final[b] = final[b], final[a]
			}
			fmt.Println("final: ", final)
		}

	var checksum int

	for i:=0; i < len(final); i++{
		n, err := strconv.Atoi(final[i])
		if err != nil{
			break
		}
		checksum += i * n
	}

	fmt.Println("checksum: ", checksum)
}