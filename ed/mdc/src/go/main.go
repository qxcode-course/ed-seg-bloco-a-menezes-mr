package main

import (
	"fmt"
)

func mdc(a, b int) int {

	rest := a % b
	if rest == 0 {
		return b
	}else {
		a = b
		b = rest
	}

	return mdc(a, b)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
