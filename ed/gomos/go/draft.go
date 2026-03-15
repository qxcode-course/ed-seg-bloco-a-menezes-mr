package main
import "fmt"
func main() {

    var d string
    var q int
    fmt.Scan(&q, &d)

    cobra := make([]int, q*2)

    for i := 0; i < q*2; i++ {
        fmt.Scan(&cobra[i])
    }

    for i := (q*2) - 1; i > 1; i-- {
        cobra[i] = cobra[i-2]
    }

    switch d {
        case "L":
            cobra[0]-- 
        case "R":
            cobra[0]++
        case "U":
            cobra[1]++
        case "D":
            cobra[1]--
    }

    for i := 0; i < q*2; i++ {
        if i % 2 == 0 {
            fmt.Printf("%d ", cobra[i])
        } else {
            fmt.Printf("%d\n", cobra[i])
        }
    }
}
