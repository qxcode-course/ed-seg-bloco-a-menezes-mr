package main

import (
	"bufio"
	"fmt"
	"os"
)

// Não mude a assinatura desta função, ela é a função chamada pelo LeetCode
func exist(grid [][]byte, word string) bool {

	for i := range len(grid) {
		for j := range len(grid[0]) {
			if search(grid, word, i, j, 0){
				return true
			}
		}
	}

	return false
}

func search(grid [][]byte, word string, i, j, aux int) bool{
	if aux == len(word) {
		return true
	}

	if i < 0 || i >= len(grid) || j < 0 || j >= len(grid[0]) {
		return false
	}

	if grid[i][j] != word[aux] {
		return false
	}

	lo := grid[i][j]
	grid[i][j] = '.'

	prox := search(grid , word, i - 1, j, aux + 1) || search(grid , word, i + 1, j, aux + 1) || search(grid , word, i, j - 1, aux + 1) || search(grid , word, i, j + 1, aux + 1)

	grid[i][j] = lo
	return prox

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var word string
	fmt.Sscanf(scanner.Text(), "%s", &word)
	grid := make([][]byte, 0)
	for scanner.Scan() {
		grid = append(grid, []byte(scanner.Text()))
	}
	if exist(grid, word) {
		fmt.Println("true")
	} else {
		fmt.Println("false")
	}
}
