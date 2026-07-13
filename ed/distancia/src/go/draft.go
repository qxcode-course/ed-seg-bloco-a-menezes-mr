package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func colocar(texto []byte, p, l int, num byte) bool {
    for i := 1; i <= l; i++ {
        esq := p - i
        if esq >= 0 && texto[esq] == num {
            return false
        }
    }

    for i := 1; i <= l; i++ {
        dir := p + i
        if dir < len(texto) && texto[dir] == num {
            return false
        }
    }

    return true
}

func resolver(texto []byte, p, l int) bool{
    if p == len(texto) {
        return true
    }
    if texto[p] != '.' {
        return resolver(texto, p + 1, l)
    }

    for i := 0; i <= l; i++ {
        num := byte('0' + i)

        if colocar(texto, p, l, num) {
            texto[p] = num
            if resolver(texto, p + 1, l) {
                return true
            }
            texto[p] = '.'
        }

    }
    return false
}

func main() {
    scanner := bufio.NewScanner(os.Stdin)
    if !scanner.Scan() {
        return
    }
    linha := scanner.Text()
    texto := []byte(linha)

    if !scanner.Scan() {
        return
    }

    l, _ := strconv.Atoi(scanner.Text())

    resolver(texto, 0, l)

    fmt.Println(string(texto))

}