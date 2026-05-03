package main

import (
	"fmt"
	"strings"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	values := []string{}
	for n := l.root.next; n != l.root; n = n.next {
		if n == sword {
			values = append(values, fmt.Sprint(n.Value,">"))
		}else {
			values = append(values, fmt.Sprint(n.Value))
		}
	}
	return "[ " + strings.Join(values, " ") + " ]"
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.Next() == l.root {
		return l.root.next
	}
	return it.Next()
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
