package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func matchingStrings(stringsList []string, queries []string) []int {

	frequencia := make(map[string]int)
	for _, s := range stringsList {
		frequencia[s]++
	}

	resultado := make([]int, len(queries))
	for i, q := range queries {
		resultado[i] = frequencia[q]
	}

	return resultado
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	scanner.Scan()
    scanner.Scan()
	stringsList := strings.Fields(scanner.Text())
	
    scanner.Scan()
    scanner.Scan()
	queries := strings.Fields(scanner.Text())

	resultados := matchingStrings(stringsList, queries)

	for i, res := range resultados {
		if i > 0 {
			fmt.Print(" ")
		}
		fmt.Print(res)
	}
	fmt.Println()
}