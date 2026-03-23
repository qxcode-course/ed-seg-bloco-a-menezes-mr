package main
import "fmt"

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

        combate[aux02] = 1
        size--;

        aux01 = (aux02 + 1) % n
        aux02 = (aux01 + 1) % n

    }
}
