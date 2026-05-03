package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

type Node struct {
	Value int
	next *Node
	prev *Node
}

type LList struct {
	root *Node
}

func NewLList() *LList{
	root := &Node{}
	root.next = root
	root.prev = root

	return &LList{
		root: root,
	}
}

func (ll *LList) Size() int {
	cont := 0
	if ll.root.next == ll.root {
		return cont
	}
	element := ll.root.next
	for i := element; i != ll.root; i = i.next {
		cont++
	}
	return cont
}

func (ll *LList) PushFront(value int){
	newNode := &Node{
		Value: value,
		next: ll.root.next,
		prev: ll.root,
	}
	newNode.next.prev = newNode
	ll.root.next = newNode
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{
		Value: value,
		next: ll.root,
		prev: ll.root.prev,
	}
	newNode.next.prev = newNode
	newNode.prev.next = newNode
}

func (ll *LList) PopBack() {
	if ll.Size() != 0 {
		ll.root.prev.prev.next = ll.root
		ll.root.prev = ll.root.prev.prev
	}
}

func (ll *LList) PopFront() {
	if ll.Size() != 0 {
		ll.root.next.next.prev = ll.root
		ll.root.next = ll.root.next.next
	}
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll *LList) String() string {
	return "[" + ll.Join(", ") + "]"
}

func (ll *LList) Join(sep string) string {
	if ll.Size() == 0 {
		return ""
	}
	element := ll.root.next
	result := fmt.Sprintf("%d", element.Value)
	for i := element.next; i != ll.root; i = i.next {
		result += sep + fmt.Sprintf("%d", i.Value)
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
// 	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			// fmt.Println(ll.String())
		case "size":
			// fmt.Println(ll.Size())
		case "push_back":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushBack(num)
			// }
		case "push_front":
			// for _, v := range args[1:] {
			// 	num, _ := strconv.Atoi(v)
			// 	ll.PushFront(num)
			// }
		case "pop_back":
			// ll.PopBack()
		case "pop_front":
			// ll.PopFront()
		case "clear":
			// ll.Clear()
		case "walk":
			// fmt.Print("[ ")
			// for node := ll.Front(); node != nil; node = node.Next() {
			// 	fmt.Printf("%v ", node.Value)
			// }
			// fmt.Print("]\n[ ")
			// for node := ll.Back(); node != nil; node = node.Prev() {
			// 	fmt.Printf("%v ", node.Value)
			// }
			// fmt.Println("]")
		case "replace":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	node.Value = newvalue
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "insert":
			// oldvalue, _ := strconv.Atoi(args[1])
			// newvalue, _ := strconv.Atoi(args[2])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Insert(node, newvalue)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "remove":
			// oldvalue, _ := strconv.Atoi(args[1])
			// node := ll.Search(oldvalue)
			// if node != nil {
			// 	ll.Remove(node)
			// } else {
			// 	fmt.Println("fail: not found")
			// }
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
