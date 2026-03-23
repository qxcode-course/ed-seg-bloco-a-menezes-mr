package main
import "fmt"

func saida(n int, faca int, combate []int) {
	fmt.Print("[ ")
	for i := 0; i < n; i++ {
		if i == faca {
			fmt.Printf("%d> ", i+1)
		} else if combate[i] == 0 { 
			fmt.Printf("%d ", i+1)
		}
	}
	fmt.Print("]\n")
}

func main() {

    var n, e int
    fmt.Scan(&n, &e)
    combate := make([]int, n)
    aux01 := e - 1
    aux02 := e % n
    size := n

    for size > 1 {
        if combate[aux01] == 1 {
            aux01 = (aux01 + 1) % n
            continue
        }
        if combate[aux02] == 1 || aux02 == aux01 {
            aux02 = (aux02 + 1) % n
            continue
        }

        saida(n, aux01, combate)

        combate[aux02] = 1
        size--;

        if size == 1 {
            saida(n, aux01, combate)
        }
        
        aux01 = (aux02 + 1) % n
        aux02 = (aux01 + 1) % n

    }
}
