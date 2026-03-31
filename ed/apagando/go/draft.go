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
	saiu := make(map[int]bool, qtd)

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

	for _, p := range filaNova {
    fmt.Print(p, " ")
  }
  fmt.Println()
}
