package main

import (
	"fmt"
	"math"
)
func main() {
    var n, a, b int
    fmt.Scan(&n)
    jogo := make([]int, n)

    for i := 0; i < n; i++ {

        fmt.Scan(&a, &b)
        if a < 10 || b < 10 {
            jogo[i] = -1
        } else  {
            jogo[i] = int(math.Abs(float64(a - b)))
        }
    }

    auxId := -1
    auxDif := 99999
    
    for id, competidor := range jogo{

        if competidor != -1 && competidor < auxDif {
            auxId = id
            auxDif = competidor
        }

    }

    if auxId == -1 {
        fmt.Println("sem ganhador")
    } else {
        fmt.Println(auxId)
    }
    
}
