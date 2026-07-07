package main

import (
	"bufio"
	"fmt"
	"os"
)

func tirarIlha(grid [][]byte, l, c int){
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid) || grid[l][c] == '0' {
		return 
	}

	grid[l][c] = '0' 
	tirarIlha(grid, l - 1, c)
	tirarIlha(grid, l + 1, c)
	tirarIlha(grid, l, c - 1)
	tirarIlha(grid, l, c + 1)

}
// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	//
	count := 0
	for l := range len(grid) {
		for c := range len(grid) {
			if grid[l][c] == '1' {
				count++
				tirarIlha(grid, l, c)
			}
		}
	}
	return count
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}
