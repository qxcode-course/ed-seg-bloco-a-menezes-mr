package main
import "fmt"
func main() {
	var n, qtd, ps int
	fmt.Scan(&n)
	fila := make([]int, n)

	for i := range fila {
    fmt.Scan(&fila[i])
  }

	fmt.Scan(&qtd)
	saiu := make(map[int]bool)

	for range qtd {
		fmt.Scan(&ps)
		saiu[ps] = true
	}

	filaNova := make([]int, 0, n - qtd)

	for _, pessoa := range fila{
		if saiu[pessoa]{
			continue
		}
		filaNova = append(filaNova, pessoa)
	} 

	for i, p := range filaNova {
    if i > 0 {
      fmt.Print(" ") 
    }
    fmt.Print(p)
  }
  fmt.Println()
}
