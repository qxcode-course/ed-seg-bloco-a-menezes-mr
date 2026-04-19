package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func processa(vet []int) {
	// fmt.Println("[ " + Join(vet, " ") + " ]")
	num := len(vet);
	if len(vet) == 1{
		return
	}
	aux := make([]int, 0 , num-1)

	for i:= 0; i < num - 1; i++ {
		aux = append(aux, vet[i] + vet[i + 1])
	}
	processa(aux)
	fmt.Println("[ " + Join(aux, " ") + " ]")

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
	fmt.Println("[ " + Join(vet, " ") + " ]")
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}
