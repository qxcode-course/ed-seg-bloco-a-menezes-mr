package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	Value int
	Left  *Node
	Right *Node
}

func BstInsert(values []int) *Node {
	
	var root *Node
	for _, elem := range values {
		if root == nil {
			root = &Node{Value: elem}
			continue
		}

		atual := root

		for {
			if elem < atual.Value {
				if atual.Left == nil {
					atual.Left = &Node{Value: elem}
					break
				}
				atual = atual.Left
			} else if elem > atual.Value {
				if atual.Right == nil {
					atual.Right = &Node{Value: elem}
					break
				}
				atual = atual.Right
			} else {
				break
			}
		}
	}
	return root
}

// Dica: crie um vetor compartilhado e vá preenchendo conforme anda na recursão
// Depois use o strings.Join para gerar o serial

func preencher(node *Node, nodes *[]string){
	if node == nil {
		*nodes = append(*nodes, "#")
		return
	}

	*nodes = append(*nodes, strconv.Itoa(node.Value))
	preencher(node.Left, nodes)
	preencher(node.Right, nodes)
}


func Serialize(root *Node) string {
	// TODO
	var nodes []string
	preencher(root, &nodes)
	return strings.Join(nodes, " ")
}

// -----------------------------------------------------------------------------------
func BShow(node *Node, history string) {
	if node != nil && (node.Left != nil || node.Right != nil) {
		BShow(node.Left, history+"l")
	}
	for i := 0; i < len(history)-1; i++ {
		if history[i] != history[i+1] {
			fmt.Print("│   ")
		} else {
			fmt.Print("    ")
		}
	}
	if history != "" {
		if history[len(history)-1] == 'l' {
			fmt.Print("╭───")
		} else {
			fmt.Print("╰───")
		}
	}
	if node == nil {
		fmt.Println("#")
		return
	}
	fmt.Println(node.Value)
	if node.Left != nil || node.Right != nil {
		BShow(node.Right, history+"r")
	}
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	nodes := strings.Split(scanner.Text(), " ")
	values := make([]int, 0, len(nodes))
	for _, elem := range nodes {
		v, err := strconv.Atoi(elem)
		if err == nil {
			values = append(values, v)
		}
	}
	root := BstInsert(values)
	BShow(root, "") // Chama a função de impressão formatada
	fmt.Println(Serialize((root)))
}
