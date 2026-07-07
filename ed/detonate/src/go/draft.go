package main

import (
	"fmt"
)

type Pos struct{
    l, c, r int
}

func detonar(atual int, visitado []bool, listaBom []Pos) int {
    visitado[atual] = true
    count := 1
    bombaA := listaBom[atual]

    for i := range len(listaBom) {
        if !visitado[i] {
            alvo := listaBom[i]

            deltL := bombaA.l - alvo.l
            deltC := bombaA.c - alvo.c
            raio := bombaA.r
            if deltL * deltL + deltC * deltC <= raio * raio {
                count += detonar(i, visitado, listaBom)
            }
        }
    }

    return count
}

func bombas(grid[][]int) int{
    n := len(grid)
    listaBom := make([]Pos, n)

    for i, b := range grid {
        listaBom[i] = Pos{l: b[0], c: b[1], r: b[2]}
    }
    detonadas := 0
    for i := range n {
        visitado := make([]bool, n)
        count := detonar(i, visitado, listaBom)

        if count > detonadas {
            detonadas = count
        }
    }
    return detonadas
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
    fmt.Println(bombas(mat))
}