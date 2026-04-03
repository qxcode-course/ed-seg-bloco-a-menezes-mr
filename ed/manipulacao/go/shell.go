package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
	"math"
	"cmp"
)

func getMen(vet []int) []int {
	men := make([]int, 0, len(vet))

	for _, pessoa := range vet {
		if pessoa > 0 {
			men = append(men, pessoa)
		}
	}
	return men
}

func getCalmWomen(vet []int) []int {
	women := make([]int, 0, len(vet))

	for _, pessoa := range vet {
		if pessoa < 0 && pessoa > -10{
			women = append(women, pessoa)
		}
	}
	return women
}

func sortVet(vet []int) []int {
	slices.Sort(vet)
	return vet
}

func sortStress(vet []int) []int {
	slices.SortFunc(vet, func(a, b int) int {
		return cmp.Compare(math.Abs(float64(a)), math.Abs(float64(b)))
	})
	return vet
}

func reverse(vet []int) []int {
	vetR := make([]int, len(vet))
	copy(vetR, vet)
	slices.Reverse(vetR)

	return vetR
}

func unique(vet []int) []int {
	elementos := make(map[int]bool)
	vetU := make([]int, 0, len(vet))

	for _, elemento := range vet {
		if !elementos[elemento] {
			vetU = append(vetU, elemento)
			elementos[elemento] = true
		}
	}

	return vetU
}

func repeated(vet []int) []int {
	elementos := make(map[int]bool)
	vetR := make([]int, 0, len(vet))

	for _, elemento := range vet {
		if elementos[elemento] {
			vetR = append(vetR, elemento)
		}else {
			elementos[elemento] = true
		}
	}

	return vetR
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}

