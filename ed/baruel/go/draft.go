package main

import (
	"fmt"
	"strconv"
)

func main() {

	var n, x, y int
	var rep string = ""
	var fal string = ""

	fmt.Scan(&n, &x)
	album := make([]int, n)

	for i := 0; i < x; i++ {
		fmt.Scan(&y)
		album[y-1]++
	}

	for i := 0; i < n; i++ {
		if album[i] != 0 {

			for j := 0; j < album[i]-1; j++ {
				if rep != "" {
					rep += " "
				}
				rep += strconv.Itoa(i+1)
			}
		} else {
			if fal != "" {
				fal += " "
			}
			fal += strconv.Itoa(i+1)
		}
	}

	if rep == "" {
		rep = "N"
	}
	if fal == "" {
		fal = "N"
	}

	fmt.Println(rep)
	fmt.Println(fal)
}