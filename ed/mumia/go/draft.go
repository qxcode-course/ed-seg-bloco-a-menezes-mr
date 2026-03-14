package main
import "fmt"
func main() {

    var nome string
    var id int

    fmt.Scan(&nome, &id)

    if id < 12 {
        fmt.Printf("%s eh crianca\n", nome)
    } else if id < 18 {
        fmt.Printf("%s eh jovem\n", nome)
    } else if id < 65 {
        fmt.Printf("%s eh adulto\n", nome)
    } else if id < 1000 {
        fmt.Printf("%s eh idoso\n", nome)
    } else {
        fmt.Printf("%s eh mumia\n", nome)
    }
    
}
