package main
import "fmt"

func resto(num int) string {
    if num == 0 {
        return ""
    }
    
    resDiv := num / 2
    restDiv := num % 2 
    resultado := resto(resDiv)
    
    return resultado + fmt.Sprintf("%d %d\n", resDiv, restDiv)
}

func main() {
    var num int
    fmt.Scan(&num)
    fmt.Print(resto(num))
}