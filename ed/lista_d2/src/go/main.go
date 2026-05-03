package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
)
// ------------- Node

type Node struct {
	Value int
	next *Node
	prev *Node
	root *Node
}

func (n *Node) Next() *Node {
	if n.next == n.root {
		return nil
	}
	return n.next
}

func (n *Node) Prev() *Node {
	if n.prev == n.root {
		return nil
	}
	return n.prev
}

//-------------- LList

type LList struct {
	root *Node
	size int
}

func NewLList() *LList{
	root := &Node{}
	root.root = root
	root.next = root
	root.prev = root

	return &LList{
		root: root,
		size: 0,
	}
}

func (ll *LList) Size() int {
	return ll.size
}

func (ll *LList) PushFront(value int){
	newNode := &Node{
		Value: value,
		next: ll.root.next,
		prev: ll.root,
		root: ll.root,
	}
	newNode.next.prev = newNode
	ll.root.next = newNode
	ll.size++
}

func (ll *LList) PushBack(value int) {
	newNode := &Node{
		Value: value,
		next: ll.root,
		prev: ll.root.prev,
		root: ll.root,
	}
	newNode.next.prev = newNode
	newNode.prev.next = newNode
	ll.size++
}

func (ll *LList) PopBack() {
	if ll.Size() != 0 {
		ll.root.prev.prev.next = ll.root
		ll.root.prev = ll.root.prev.prev
		ll.size--
	}
}

func (ll *LList) PopFront() {
	if ll.Size() != 0 {
		ll.root.next.next.prev = ll.root
		ll.root.next = ll.root.next.next
		ll.size--
	}
}

func (ll *LList) Clear() {
	ll.root.next = ll.root
	ll.root.prev = ll.root
}

func (ll *LList) Front() *Node{
	return ll.root.Next()
}

func (ll *LList) Back() *Node{
	return ll.root.Prev()
}

func (ll *LList) Search(value int) *Node{
	element := ll.root.Next()

	for i := element; i != nil ; i = i.Next() {
		if i.Value == value {
			return i
		}
	}
	return nil
}

func (ll *LList) Insert(node *Node, value int) {
	newNode := &Node{
		Value: value,
		next: node,
		prev: node.prev,
		root: node.root,
	}

	node.prev.next = newNode
	node.prev = newNode
	
	ll.size++
}

func (ll *LList) Remove(node *Node) *Node{
	if node.Next() == nil {
		return nil
	}
	node.prev.next = node.next
	node.next.prev = node.prev

	ll.size--

	return node.Next()
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
	for i := element.next; i != i.root; i = i.next {
		result += sep + fmt.Sprintf("%d", i.Value)
	}
	return result
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

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
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != nil; node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != nil; node = node.Prev() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
