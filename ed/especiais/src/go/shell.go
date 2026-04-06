package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
	"cmp"
	"math"
)

type Pair struct {
	One int
	Two int
}

func occurr(vet []int) []Pair {
	occur := make([]Pair, 0, len(vet))
	elementos := make(map[int]int)
	
	for _, elemento := range vet {
		elemento = int(math.Abs(float64(elemento)))
		if elementos[elemento] == 0 {
			elementos[elemento] = 1
		}else {
			elementos[elemento] += 1
		}
	}
	for key, valor := range elementos {
		occur = append(occur, Pair{One: key, Two: valor})
	}

	slices.SortFunc(occur, func(a, b Pair) int { return cmp.Compare(a.One, b.One)})
	return occur
}

func teams(vet []int) []Pair {
	if len(vet) == 0 {
		return []Pair{} 
	}

	team := make([]Pair, 0, len(vet))
	valor := vet[0]
	rep := 1

	for i := 1; i < len(vet); i++ {

		if vet[i] == valor {
			rep++

		} else {
			team = append(team, Pair{One: valor, Two: rep})
			valor = vet[i]
			rep = 1
		}
	}

	team = append(team, Pair{One: valor, Two: rep})
	return team
}

func mnext(vet []int) []int {
	mapa := make([]int,len(vet))

	for i := 0; i < len(vet); i++ {
		me := i > 0 && vet[i-1] < 0
		md := i < len(vet) - 1 && vet[i+1] < 0

		if me || md {
			mapa[i] = 1
		}
	}

	return mapa
}

func alone(vet []int) []int {
	mapa := make([]int,len(vet))
	if len(vet) == 1 && vet[0] > 0 {
		mapa[0] = 1
		return mapa
	}

	for i := 0; i < len(vet); i++ {
		me := i > 0 && vet[i-1] < 0
		md := i < len(vet) - 1 && vet[i+1] < 0

		if !(me || md) {
			mapa[i] = 1
		}
	}

	return mapa
}

func couple(vet []int) int {
	
	solt := make(map[int]int)
	cont := 0
	for _, elemento := range vet {
		if solt[-elemento] > 0 {
			cont ++
			solt[-elemento]--
		}else {
			solt[elemento]++
		}
	}

	return cont
}

func hasSubseq(vet []int, seq []int, pos int) bool {
	_ = vet
	_ = seq
	_ = pos
	return false
}

func subseq(vet []int, seq []int) int {
	cont := 0
	id := 0
	for i, elemento := range vet {
		if elemento == seq[cont] {
			cont++
			if cont == len(seq) {
				return id
			}
		}else if elemento == seq[0] {
			cont = 1
			id = i
			if cont == len(seq) {
				return id
			}
		}else {
			id = i+1
			cont = 0
		}
	}
	return -1
}

func erase(vet []int, posList []int) []int {

	mapa := make(map[int]bool, len(posList))
	size := len(vet) - len(posList)
	newList := make([]int, 0, size)

	if size == 0 { return newList}

	for _, id := range posList {
		mapa[id] = true
	}
	
	for i, elemento := range vet {
		if !mapa[i] {
			newList = append(newList, elemento)
		}
	}

	return newList
}

func clear(vet []int, value int) []int {

	newList := make([]int, 0, len(vet))
	
	for _, elemento := range vet {
		if elemento != value {
			newList = append(newList, elemento)
		}
		
	}
	return newList
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "occurr":
			printSlice(occurr(str2vet(args[1])))
		case "teams":
			printSlice(teams(str2vet(args[1])))
		case "mnext":
			printSlice(mnext(str2vet(args[1])))
		case "alone":
			printSlice(alone(str2vet(args[1])))
		case "erase":
			printSlice(erase(str2vet(args[1]), str2vet(args[2])))
		case "clear":
			val, _ := strconv.Atoi(args[2])
			printSlice(clear(str2vet(args[1]), val))
		case "subseq":
			fmt.Println(subseq(str2vet(args[1]), str2vet(args[2])))
		case "couple":
			fmt.Println(couple(str2vet(args[1])))
		case "end":
			return
		default:
			fmt.Println("Invalid command")
		}
	}
}

// Funções auxiliares

func str2vet(str string) []int {
	if str == "[]" {
		return nil
	}
	str = str[1 : len(str)-1]
	parts := strings.Split(str, ",")
	var vet []int
	for _, part := range parts {
		num, _ := strconv.Atoi(strings.TrimSpace(part))
		vet = append(vet, num)
	}
	return vet
}

func printSlice[T any](vet []T) {
	fmt.Print("[")
	for i, x := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(x)
	}
	fmt.Println("]")
}

func (p Pair) String() string {
	return fmt.Sprintf("(%v, %v)", p.One, p.Two)
}
