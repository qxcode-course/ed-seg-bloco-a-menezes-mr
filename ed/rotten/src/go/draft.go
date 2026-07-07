package main

import (
	"fmt"
)

type Pos struct{
    l, c int
}

func laranjaP(grid[][]int, l, c int) int{
    fila := []Pos {}
    count := 0

    for i := range l {
        for j := range c {
            if grid[i][j] == 2 {
                fila = append(fila, Pos{i, j})
            }else if grid[i][j] == 1 {
                count++
            }
        }
    }

    if count == 0 { return 0}

    minuto := 0
    direcao := []Pos {{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
    for len(fila)  > 0 && count > 0{
        size := len(fila)

        for range size {
            atual := fila[0]
            fila = fila[1:]

            for _, d := range direcao {
                nl := atual.l+d.l
                nc := atual.c+d.c
                if nl >= 0 && nl < l && nc >= 0 && nc < c && grid[nl][nc] == 1 {
                    grid[nl][nc] = 2
                    count--
                    fila = append(fila, Pos{nl, nc})
                }
            }
        }
        minuto++

    }
    if count > 0 {
        return -1
    }

    return minuto

}



func main() {
	var m, n int
	fmt.Scan(&m, &n)
	mat := make([][]int, m)

    for i := range m {
        mat[i] = make([]int, n)
        for j := range n {
            fmt.Scan(&mat[i][j])
        }
        
    }
    fmt.Println(laranjaP(mat, m, n))
}