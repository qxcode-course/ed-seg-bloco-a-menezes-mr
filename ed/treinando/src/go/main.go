package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func join(vet []int, aux string) string{
	if len(vet) == 1 {
		str := fmt.Sprint(vet[0])
		return str
	}

	str := fmt.Sprint(vet[0],aux)
	return str + join(vet[1:], aux)
}


func tostr(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	return ("[" + join(vet, ", ") + "]")
}

func tostrrev(vet []int) string {
	if len(vet) == 0 {
		return "[]"
	}
	reverse(vet)
	return ("[" + join(vet, ", ") + "]")
	// return ""
}
// reverse: inverte os elementos do slice
func reverse(vet []int) {
	tam := len(vet)
	if tam == 1 || tam == 0 {
		return
	}

	i := vet[0]
	j := vet[tam - 1]

	vet[0] = j
	vet[tam - 1] = i

	reverse(vet[1:tam-1])
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	tam := len(vet)
	if tam == 0 {
		return 0
	}
	if tam == 1{
		return vet[0]
	}
	return vet[0] + sum(vet[1:])
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	tam := len(vet)
	if tam == 0 {
		return 1
	}
	if tam == 1{
		return vet[0]
	}
	return vet[0] * mult(vet[1:])
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return -1
	}
	id, value := 0, vet[0]
	cont := id
	var rec func(v []int) (int, int)
	rec = func(v []int) (int, int){
		if len(v) == 0 {
			return value, id
		}
		if value > v[0] {
			value = v[0]
			id = cont
		}
		cont ++
		return rec(v[1:])
	}

	rec(vet)
	return id
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
