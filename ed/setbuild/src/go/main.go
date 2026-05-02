package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)

type Vector struct{
	data []int
	size int
	capacity int
}

func NewSet(capacity int) *Vector{
	return &Vector{
		data: make([]int, capacity),
		size: 0,
		capacity: capacity,

	}
}
func (v *Vector) ToString() string { 
	if v.size == 0 {
		return "[]"
	}
	return "[" + Join(v.data[:v.size], ", ") + "]"
}

func (v *Vector) reserve(newCapacity int) {
	newData := make([]int, newCapacity)
	copy(newData, v.data[:v.size])
	v.data = newData
	v.capacity = newCapacity
}

func (v *Vector) binarySearch(value int) int {
	inicio := 0
	fim := v.size - 1

	for inicio <= fim {		
		meio := inicio + (fim - inicio)/2

		if v.data[meio] == value {
			return meio
		}
		if v.data[meio] < value {
			inicio = meio + 1
		} else {
			fim = meio - 1
		}
	}
	return -1
}
func (v *Vector) Contains(value int) bool {
	flag := v.binarySearch(value) 
	if flag == -1{
		return false
	}
	return true
}


func (v *Vector) insert(value int, index int) error {
	if v.size == v.capacity {
		v.reserve(v.capacity * 2)
	}
	
	for i:= v.size; i > index; i-- {
		v.data[i] = v.data[i-1];
	}
	v.data[index] = value
	v.size ++
	return nil
}


func (v *Vector) Insert(value int){
	if !v.Contains(value) {
		inicio := 0
		fim := v.size - 1

		for inicio <= fim {		
			meio := inicio + (fim - inicio)/2
			if v.data[meio] < value {
				inicio = meio + 1
			} else {
				fim = meio - 1
			}
		}
		v.insert(value, inicio)
	}

}

func Join(slice []int, sep string) string {
	if len(slice) == 0 {
		return ""
	}
	var result strings.Builder
	fmt.Fprintf(&result, "%d", slice[0])
	for _, value := range slice[1:] {
		fmt.Fprintf(&result, "%s%d", sep, value)
	}
	return result.String()
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.ToString())
		case "erase":
			// value, _ := strconv.Atoi(parts[1])
		case "contains":
			// value, _ := strconv.Atoi(parts[1])
		case "clear":
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
