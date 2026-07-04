package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func burnTrees(grid [][]rune, l, c int) {
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})

	auxl := []int {-1, 1, 0, 0}
	auxc := []int {0, 0, -1, 1}
	for !stack.IsEmpty(){
		elem := stack.Pop()
		if elem.l >= 0 && elem.l < len(grid) && elem.c >= 0 && elem.c < len(grid[0]){
			if grid[elem.l][elem.c] == '#' {
				grid[elem.l][elem.c] = 'o'

				for i := range 4 {
					vl := elem.l + auxl[i]
					vc := elem.c + auxc[i]

					stack.Push(Pos{vl, vc})
				}
			}
		}
	}


}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
