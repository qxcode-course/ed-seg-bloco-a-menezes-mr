package main

import (
	"bufio"
	"fmt"
	"os"
)

func marcar(grid [][]byte, l, c int) {
	if l < 0 || l >= len(grid) || c < 0 || c >= len(grid[0]) || grid[l][c] != 'O' {
		return
	}

	grid[l][c] = '#'

	marcar(grid, l - 1, c)
	marcar(grid, l + 1, c)
	marcar(grid, l, c - 1)
	marcar(grid, l, c + 1)

}
// NÃO ALTERE A ASSINATURA DA FUNÇÃO solve
func solve(board [][]byte) {
	nl := len(board)
	nc := len(board[0])

	for i := range nl {
		if board[i][0] == 'O' {
			marcar(board, i, 0)
		}

		if board[i][nc - 1] == 'O' {
			marcar(board, i, nc - 1)
		}
	}

	for i := range nc {
		if board[0][i] == 'O' {
			marcar(board, 0, i)
		}

		if board[nl - 1][i] == 'O' {
			marcar(board, nl - 1, i)
		}
	}

	for i := range nl {
		for j := range nc {
			if board[i][j] == 'O' {
				board[i][j] = 'X'
			}else if board[i][j] == '#' {
				board[i][j] = 'O'
			}
		}
	}

}

// NÃO ALTERE A MAIN
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var nrows, ncols int
	fmt.Sscanf(scanner.Text(), "%d %d", &nrows, &ncols)
	board := make([][]byte, nrows)
	for i := 0; i < nrows; i++ {
		scanner.Scan()
		board[i] = []byte(scanner.Text())
	}
	solve(board)
	for _, row := range board {
		fmt.Println(string(row))
	}
}
