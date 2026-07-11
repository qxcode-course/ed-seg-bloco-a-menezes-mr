package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func memorisar(matrix[][]int, m [][]int, l, c int) int{
	if m[l][c] != 0 {
		return m[l][c]
	}

	caminhoM := 1
	dirL:= []int {-1, 1, 0, 0}
	dirC:= []int {0, 0, -1, 1}

	for i := range 4 {
		visL := l + dirL[i]
		visC := c + dirC[i]

		if visL >= 0 && visL < len(matrix) && visC >= 0 && visC < len(matrix[0]) && matrix[visL][visC] > matrix[l][c] {
			caminhoV := 1 + memorisar(matrix, m, visL, visC)

			if caminhoV > caminhoM {
				caminhoM = caminhoV
			}
		}

	}

	m[l][c] = caminhoM
	return caminhoM
}

func longestIncreasingPath(matrix [][]int) int {
	//
	l := len(matrix)
	c := len(matrix[0])

	m := make([][]int, l)
	for i := range l {
		m[i] = make([]int, c)
	}

	recorde := 0

	for i := range l {
		for j := range c {
			caminhoA := memorisar(matrix, m, i, j)

			if caminhoA > recorde {
				recorde = caminhoA
			}
		}
	}

	return recorde
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)

	if !scanner.Scan() {
		return
	}
	parts := strings.Fields(scanner.Text())
	if len(parts) < 2 {
		return
	}
	nl, _ := strconv.Atoi(parts[0])
	nc, _ := strconv.Atoi(parts[1])

	matrix := make([][]int, nl)
	for i := 0; i < nl; i++ {
		if !scanner.Scan() {
			return
		}
		tokens := strings.Fields(scanner.Text())
		row := make([]int, nc)
		for j := 0; j < nc && j < len(tokens); j++ {
			v, _ := strconv.Atoi(tokens[j])
			row[j] = v
		}
		matrix[i] = row
	}

	fmt.Println(longestIncreasingPath(matrix))
}
