package main

import "fmt"

func main() {

    var n, t int
    fmt.Scan(&n)
    casais := make([]int, n)

    for i := 0; i < n; i++ {
        var a int
        fmt.Scan(&a)
        casais[i] = a
    }

    for i := 0; i < n; i++ {
        if casais[i] == 0 {
            continue
        }
        for j := i+1; j < n; j++ {
            if casais[i] == -casais[j]{
                t++
                casais[j] = 0
                casais[j] = 0
                break
            }
        }
    }

    fmt.Println(t)
}
