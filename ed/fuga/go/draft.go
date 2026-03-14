package main

import "fmt"

func main() {
    var h, p, f, d int
    fmt.Scan(&h, &p, &f, &d)
    
    for {

        if f == h {
            fmt.Println("S")
            break
        }
        if f == h {
            fmt.Println("N")
            break
        }

        f = ((f+d)%16 + 16) % 16
    }
}