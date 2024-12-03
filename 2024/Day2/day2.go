//Autore: Francesco Corrado

package main

import (
	"fmt"
	"bufio"
	"os"
	"strings"
	"strconv"
)

func Abs(x int) int{
	if x < 0 {
		return -x
	}
	return x
}


func main(){
	totalSafe := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		check := true
		removed := false
		var num []int
		var trend = ""
		line := strings.Split(scanner.Text(), " ")
		for i:=0; i<len(line); i++{
			x, _ := strconv.Atoi(line[i])
			num = append(num, x)
		}
		for i:=1; i<len(num); i++{
			//fmt.Println(num[i])
			//fmt.Println("check between ", num[i], num[i-1])
			if i == 1{
				if num[0] > num[1]{
					trend = "d"
				}else{
					trend = "u"
				}
			}
			if trend == "d" && num[i] > num[i-1]{
				check = false
			}else if trend == "u" && num[i] < num[i-1]{
				check = false
			}
			if Abs(num[i]-num[i-1]) < 1 || Abs(num[i]-num[i-1]) > 3{
				check = false
			}
			if check == false && removed == false{
				fmt.Println("Check if we can remove error ", num[i], "in", num)
				if i == 1{
					if num[0] > num[2]{
						trend = "d"
					}else{
						trend = "u"
					}
				}
				if i+1 > len(num)-1{
					num = num[:i]
					check = true
					break
				}

				//fmt.Println("NUMERI (", num[i-1], num[i+1] ,")")


				if trend == "d" && num[i-1] > num[i+1]{
					check = true
				}else if trend == "u" && num[i-1] < num[i+1]{
					check = true
				}
				if Abs(num[i-1]-num[i+1]) > 0 && Abs(num[i-1]-num[i+1]) < 4{
					check = true
				}


				if check == true{
					removed = true
					//fmt.Println("Removed ", num[i])
					num = append(num[:i], num[i+1:]...)
					i--
				}
			}
			if check == false{
				break
			}
		}
		fmt.Println("row: ", num, "is safe: ", check)
		if check{
			totalSafe++
		}
	}
	fmt.Println("Total Safe Rows: ", totalSafe)
}

/*
if check == false && removed == false{
				fmt.Println("Check if we can remove ", num[i])
				if i == 1{
					if num[0] > num[2]{
						trend = "d"
					}else{
						trend = "u"
					}
				}
				if i+1 > len(num)-1{
					check = true
					break
				}
				fmt.Println("NUMERI (", num[i], num[i-2] ,")")
				if trend == "d" && num[i+1] > num[i-1]{
					check = true
					removed = true
				}else if trend == "u" && num[i] < num[i-2]{
					check = true
					removed = true
				}
				if Abs(num[i]-num[i-2]) < 1 || Abs(num[i]-num[i-2]) > 3{
					check = true
					removed = true
				}
				if check == true{
					fmt.Println("Removed ", num[i])
					fmt.Println(num[:i], num[i+1:])
					num = append(num[:i], num[i+1:]...)
					i--
					//l--
				}
			}
			*/




/* Part 1
func main(){
	totalSafe := 0
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		check := true
		var num []int
		var trend = ""
		line := strings.Split(scanner.Text(), " ")
		for i:=0; i<len(line); i++{
			x, _ := strconv.Atoi(line[i])
			num = append(num, x)
		}
		for i:=1; i<len(num); i++{
			if i == 1{
				if num[0] > num[1]{
					trend = "d"
				}else{
					trend = "u"
				}
			}
			if trend == "d" && num[i] > num[i-1]{
				check = false
				break
			}else if trend == "u" && num[i] < num[i-1]{
				check = false
				break
			}
			if Abs(num[i]-num[i-1]) < 1 || Abs(num[i]-num[i-1]) > 3{
				check = false
				break
			}
		}
		if check{
			totalSafe++
		}
	}
	fmt.Println("Total Safe Rows: ", totalSafe)
}
*/