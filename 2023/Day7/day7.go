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

func fiveOfAKind(s string) bool{
	chars := strings.Split(s, "")
	for i:=1; i < len(chars); i++{
		if s[i] != s[0]{
			return false
		}
	}
	return true
}

func fourOfAKind(s string) bool{
	chars := strings.Split(s, "")
	m := make(map[string]int)
	for i:=0; i < len(chars); i++{
		m[string(chars[i])]++
	}
	for key, _ := range m {
		if m[key] == 4{
			return true
		}
	}
	return false
}

func fullHouse(s string) bool{
	chars := strings.Split(s, "")
	check1 := false
	check2 := false
	m := make(map[string]int)
	for i:=0; i < len(chars); i++{
		m[string(chars[i])]++
	}
	for key, _ := range m {
		if m[key] == 3{
			check1 = true
		}
		if m[key] == 2{
			check2 = true
		}
	}
	if check1 && check2{
		return true
	}
	return false
}

func threeOfAKind(s string) bool{
	chars := strings.Split(s, "")
	m := make(map[string]int)
	for i:=0; i < len(chars); i++{
		m[string(chars[i])]++
	}
	for key, _ := range m {
		if m[key] == 3{
			return true
		}
	}
	return false
}

func twoPair(s string) bool{
	chars := strings.Split(s, "")
	check := 0
	m := make(map[string]int)
	for i:=0; i < len(chars); i++{
		m[string(chars[i])]++
	}
	for key, _ := range m {
		if m[key] == 2{
			check++
		}
	}
	if check == 2{
		return true
	}
	return false
}

func onePair(s string) bool{
	chars := strings.Split(s, "")
	m := make(map[string]int)
	for i:=0; i < len(chars); i++{
		m[string(chars[i])]++
	}
	for key, _ := range m {
		if m[key] == 2{
			return true
		}
	}
	return false
}

func cardValue(c byte) int {
	switch c {
	case '2':
		return 2
	case '3':
		return 3
	case '4':
		return 4
	case '5':
		return 5
	case '6':
		return 6
	case '7':
		return 7
	case '8':
		return 8
	case '9':
		return 9
	case 'T':
		return 10
	case 'J':
		return 11
	case 'Q':
		return 12
	case 'K':
		return 13
	case 'A':
		return 14
	default:
		return 0
	}
}

func compareHands(a, b string) bool {
	for i := 0; i < len(a); i++ {
		if cardValue(a[i]) != cardValue(b[i]) {
			return cardValue(a[i]) > cardValue(b[i])
		}
	}
	return false
}

func checkCardQueue(cards []string) []string {
	var totalQueue []string
	var fiveQueue []string
	var fourQueue []string
	var fullQueue []string
	var threeQueue []string
	var twoQueue []string
	var oneQueue []string
	var highQueue []string

	fmt.Println("1:", cards)
	for i:=0; i < len(cards); i++{
		if fiveOfAKind(cards[i]){
			fmt.Println(cards[i], "è fiveOfAKind")
			fiveQueue = append(fiveQueue, cards[i])
			cards = append(cards[:i], cards[i+1:]...)
			i--
		}
	}

	sort.Slice(fiveQueue, func(i, j int) bool {
		return compareHands(fiveQueue[i], fiveQueue[j])
	})
	totalQueue = append(totalQueue, fiveQueue...)

	fmt.Println("coda locale: ", fiveQueue)
	fmt.Println("coda totale: ", totalQueue)


	fmt.Println("2:", cards)
	for i:=0; i < len(cards); i++{
		if fourOfAKind(cards[i]){
			fmt.Println(cards[i], "è fourOfAKind")
			fourQueue = append(fourQueue, cards[i])
			cards = append(cards[:i], cards[i+1:]...)
			i--
		}
	}
	sort.Slice(fourQueue, func(i, j int) bool {
		return compareHands(fourQueue[i], fourQueue[j])
	})
	totalQueue = append(totalQueue, fourQueue...)

	fmt.Println("coda locale: ", fourQueue)
	fmt.Println("coda totale: ", totalQueue)


	fmt.Println("3:", cards)
	for i:=0; i < len(cards); i++{
		if fullHouse(cards[i]){
			fmt.Println(cards[i], "è fullHouse")
			fullQueue = append(fullQueue, cards[i])
			cards = append(cards[:i], cards[i+1:]...)
			i--
		}
	}
	sort.Slice(fullQueue, func(i, j int) bool {
		return compareHands(fullQueue[i], fullQueue[j])
	})
	totalQueue = append(totalQueue, fullQueue...)

	fmt.Println("coda locale: ", fullQueue)
	fmt.Println("coda totale: ", totalQueue)


	fmt.Println("4:", cards)
	for i:=0; i < len(cards); i++{
		if threeOfAKind(cards[i]){
			fmt.Println(cards[i], "è threeOfAKind")
			threeQueue = append(threeQueue, cards[i])
			cards = append(cards[:i], cards[i+1:]...)
			i--
		}
	}
	sort.Slice(threeQueue, func(i, j int) bool {
		return compareHands(threeQueue[i], threeQueue[j])
	})
	totalQueue = append(totalQueue, threeQueue...)

	fmt.Println("coda locale: ", threeQueue)
	fmt.Println("coda totale: ", totalQueue)


	fmt.Println("5:", cards)
	for i:=0; i < len(cards); i++{
		if twoPair(cards[i]){
			fmt.Println(cards[i], "è twoPair")
			twoQueue = append(twoQueue, cards[i])
			cards = append(cards[:i], cards[i+1:]...)
			i--
		}
	}
	sort.Slice(twoQueue, func(i, j int) bool {
		return compareHands(twoQueue[i], twoQueue[j])
	})
	totalQueue = append(totalQueue, twoQueue...)

	fmt.Println("coda locale: ", twoQueue)
	fmt.Println("coda totale: ", totalQueue)

	fmt.Println("6:", cards)
	for i:=0; i < len(cards); i++{
		if onePair(cards[i]){
			fmt.Println(cards[i], "è onePair")
			oneQueue = append(oneQueue, cards[i])
			cards = append(cards[:i], cards[i+1:]...)
			i--
		}
	}
	sort.Slice(oneQueue, func(i, j int) bool {
		return compareHands(oneQueue[i], oneQueue[j])
	})
	totalQueue = append(totalQueue, oneQueue...)

	fmt.Println("coda locale: ", oneQueue)
	fmt.Println("coda totale: ", totalQueue)


	fmt.Println("7:", cards)
	for i:=0; i < len(cards); i++{
		fmt.Println(cards[i], "è highCard")
		highQueue = append(highQueue, cards[i])
	}
	sort.Slice(highQueue, func(i, j int) bool {
		return compareHands(highQueue[i], highQueue[j])
	})
	totalQueue = append(totalQueue, highQueue...)

	fmt.Println("coda locale: ", highQueue)
	fmt.Println("coda totale: ", totalQueue)

	return totalQueue
}

func main(){
	var hands []string
	tot := 0
	bids := make(map[string]int)
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan(){
		line := scanner.Text()
		val := strings.Split(line, " ")
		//fmt.Println(val)
		hands = append(hands, val[0])
		n, _ := strconv.Atoi(val[1])
		bids[val[0]] = n
	}
	fmt.Println(hands)
	hands = checkCardQueue(hands)

	for i:=0; i < len(hands); i++{
		v := bids[hands[i]]
		fmt.Println("calcolo: ", v, "*", len(hands)-i, "=", v*(len(hands)-i))
		tot += v * (len(hands)-i)
	}
	fmt.Println("SOMMA FINALE: ", tot)

	
	/*
	fmt.Println()
	keys := make([]string, 0, len(bids))
    for k := range bids{
        keys = append(keys, k)
    }
    sort.Strings(keys)
    for _, k := range keys {
        fmt.Println(k, "ha come bid", bids[k])
    } */
}