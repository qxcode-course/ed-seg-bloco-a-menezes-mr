package main

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Vector struct {
	data     []int
	size     int
	capacity int
}

func NewVector(capacity int) *Vector {
	return &Vector{
		data:     make([]int, capacity), // nunca use len(data) ou cap(data) ou qq método do go de manipulação de array
		size:     0,
		capacity: capacity,
	}
}

func (v *Vector) Status() string {
	return fmt.Sprintf("size:%d capacity:%d", v.size, v.capacity)
}

func (v *Vector) ToString() string { 
	if v.size == 0 {
		return "[]"
	}
	return "[" + Join(v.data[:v.size], ", ") + "]"
}

func (v *Vector) Reserve(newCapacity int){ 
	aux := make([]int, newCapacity)
	for i, elemento := range v.data {
		aux[i] = elemento
	}
	v.data = aux
	v.capacity = newCapacity
}

func (v *Vector) PushBack(value int){ 
	if v.size == v.capacity {
		v.Reserve(v.capacity * 2)
	}

	v.data[v.size] = value
	v.size++
	
}

func (v *Vector) Get(id int) int{ 
	return v.data[id]
}

func (v *Vector) At(id int) (int, error){ 
	if id >= v.size {
		return -1, errors.New("index out of range")
	}
	return v.Get(id), nil
}

func (v *Vector) Set(id int, value int) error{ 
	if id >= v.size {
		return errors.New("index out of range")
	}
	v.data[id] = value
	return nil
}

func (v *Vector) Clear(){ 
	v.size = 0
}

func (v *Vector) PopBack() (int, error){ 
	if v.size == 0 {
		return -1, errors.New("vector is empty")
	}
	elemento := v.size - 1
	v.size--
	return v.Get(elemento), nil
}

func (v *Vector) Insert(id int, value int) error{ 
	if v.size == v.capacity {
		v.Reserve(v.capacity * 2)
		// return errors.New("index out of range")
	}
	v.size++
	for i := v.size - 1; i > id; i -- {
		v.data[i] = v.data[i - 1]
	} 
	v.data[id] = value

	return nil
}

func (v *Vector) Erase(id int) error{ 
	_, err := v.At(id)
	if err != nil {
		return errors.New("index out of range")
	}
	v.size--
	for i := id; i < v.size; i ++ {
		v.data[i] = v.data[i + 1]
	}
	return nil
}

func (v *Vector) IndexOf(value int) int{ 
	for i, elemento := range v.data {
		if elemento == value{
			return i
		}
	}
	return -1
}

func (v *Vector) Contains(value int) bool{ 
	idTrue := v.IndexOf(value)
	if idTrue == -1 {
		return false
	}
	return true
}

func (v *Vector) Slice(start int, end int) *Vector {
	if end < 0 {
		end += v.size
	}
	novoSize := end - start
	novoData := v.data[start:end]

	return &Vector{
		data: novoData,
		size: novoSize,
		capacity: v.capacity,
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

	v := NewVector(0)
	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
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
			v = NewVector(value)
		case "push":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.PushBack(value)
			}
		case "show":
			fmt.Println(v.ToString())
		case "status":
			fmt.Println(v.Status())
		case "pop":
			_, err := v.PopBack()
			if err != nil {
				fmt.Println(err)
			}
		case "insert":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Insert(index, value)
			if err != nil {
				fmt.Println(err)
			}
		case "erase":
			index, _ := strconv.Atoi(parts[1])
			err := v.Erase(index)
			if err != nil {
				fmt.Println(err)
			}
		case "indexOf":
			value, _ := strconv.Atoi(parts[1])
			index := v.IndexOf(value)
			fmt.Println(index)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if v.Contains(value) {
				fmt.Println("true")
			} else {
				fmt.Println("false")
			}
		case "clear":
			v.Clear()
		case "capacity":
			// fmt.Println(v.Capacity())
		case "get":
			index, _ := strconv.Atoi(parts[1])
			value, err := v.At(index)
			if err != nil {
				fmt.Println(err)
			} else {
				fmt.Println(value)
			}
		case "set":
			index, _ := strconv.Atoi(parts[1])
			value, _ := strconv.Atoi(parts[2])
			err := v.Set(index, value)
			if err != nil {
				fmt.Println(err)
			}
			
		case "reserve":
			newCapacity, _ := strconv.Atoi(parts[1])
			v.Reserve(newCapacity)
		case "slice":
			start, _ := strconv.Atoi(parts[1])
			end, _ := strconv.Atoi(parts[2])
			slice := v.Slice(start, end)
			fmt.Println(slice.ToString())
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
