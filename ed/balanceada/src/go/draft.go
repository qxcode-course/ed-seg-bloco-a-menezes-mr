package main

import (
	"bufio"
	"fmt"
	"os"
)

func balanceado(s string) bool {
    var pilha []rune
    for _, c := range s {
        if c == '('|| c == '[' {
            pilha = append(pilha, c)
        } else if c == ')'|| c == ']' {
            if len(pilha) == 0 {
                return false
            }

            topo := pilha[len(pilha) - 1]
            pilha = pilha[:len(pilha) - 1]
            if (c == ')' && topo != '(') || (c == ']' && topo != '[') {
                return false
            }
        }
    }
    return len(pilha) == 0
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()

    if balanceado(line) {
        fmt.Println("balanceado")
    } else {
        fmt.Println("nao balanceado")
    }
}