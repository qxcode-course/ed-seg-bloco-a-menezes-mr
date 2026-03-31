package main
import "fmt"
func main() {
	var n, qtd, ps int
	fmt.Scan(&n)
	fila := make([]int, n)

	for _, pessoa := range fila{
		fmt.Scan(&pessoa)
	}

	fmt.Scan(&qtd)
	saiu := make(map[int]bool)

	for range qtd {
		fmt.Scan(&ps)
		saiu[ps] = true
	}

	filaNova := make([]int, n - qtd)
	for i := 0; i < n - qtd; i++ {

	}

}
