package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MultiSet struct {
	data []int
	size int
	capacity int
}

func NewMultiSet(capacity int) *MultiSet {
	return &MultiSet{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,
	}
}

func (ms *MultiSet) expand(){
	if ms.capacity == 0 {
		ms.capacity = 1
	}else {
		ms.capacity *= 2
	}
	aux := make([]int, ms.capacity)
	copy(aux, ms.data[:ms.size])
	ms.data = aux

}

func (ms *MultiSet) search(value int) (bool,int) {
	inicio := 0
	fim := ms.size - 1

	for inicio <= fim {
		meio := inicio + (fim - inicio)/2

		if ms.data[meio] == value {
			return true, meio
		}

		if ms.data[meio] < value {
			inicio = meio + 1
		}else {
			fim = meio - 1
		}
	}

	return false, inicio
}

func (ms *MultiSet) Contains(value int) bool {
	existe , _ := ms.search(value)
	return existe
}
func (ms *MultiSet) insert(value int, index int) {
	if ms.size == ms.capacity {
		ms.expand()
	}
	for i:= ms.size; i > index; i-- {
		ms.data[i] = ms.data[i-1]
	}
	ms.data[index] = value
	ms.size++
}

func (ms *MultiSet) Insert(value int) {
	existe, i := ms.search(value)
	if existe {
		ms.insert(value, i+1)
	} else{
		ms.insert(value, i)
	}
}

func (ms *MultiSet) String() string{
	if ms.size == 0 {
		return "[]"
	}
	return "[" + Join(ms.data[:ms.size], ", ") + "]"
}

func (ms *MultiSet) Erase(value int) bool {
	existe, index := ms.search(value)
	if !existe{
		return false
	}
	for i:= index; i < ms.size - 1; i++ {
		ms.data[i] = ms.data[i+1];
	}
	ms.size --
	return true
}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	result := fmt.Sprintf("%d", slice[0])
	for _, value := range slice[1:] {
		result += sep + fmt.Sprintf("%d", value)
	}
	return result
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)
	ms := NewMultiSet(0)

	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		args := strings.Fields(line)
		fmt.Println(line)
		if len(args) == 0 {
			continue
		}
		cmd = args[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(args[1])
			ms = NewMultiSet(value)
		case "insert":
			for _, part := range args[1:] {
				value, _ := strconv.Atoi(part)
				ms.Insert(value)
			}
		case "show":
			fmt.Println(ms.String())
		case "erase":
			value, _ := strconv.Atoi(args[1])
			if !ms.Erase(value) {
				fmt.Println("value not found")
			}
		case "contains":
			value, _ := strconv.Atoi(args[1])
			fmt.Println(ms.Contains(value))
		case "count":
			// value, _ := strconv.Atoi(args[1])
		case "unique":
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
