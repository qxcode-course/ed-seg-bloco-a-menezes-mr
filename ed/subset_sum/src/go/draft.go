package main

import (
	"fmt"
)

func search(list []int, j, i int) bool{
    if j == 0 { return true }
    if i == len(list) || j < 0 { return false}
    inList := search(list, j - list[i], i +1)

    if inList { return true}
    notList := search(list, j, i +1)
    return notList
}

func main() {
    n, k := 0, 0
    fmt.Scan(&n, &k)

    nums := make([]int,n)
    for i := range n {
        fmt.Scan(&nums[i])
    }

    if search(nums, k, 0) {
        fmt.Println("true")
    } else {
        fmt.Println("false")
    }

}