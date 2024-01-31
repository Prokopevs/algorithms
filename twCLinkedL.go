package main 

import (
	"fmt"
)


type Node struct {
	data string
	next *Node
	prev *Node
}

var head *Node = new(Node)
var tail *Node = new(Node)

func main() {
	head.data = "A"
	head.next = nil
	head.prev = nil

	B := &Node{data: "B", next: nil, prev: head}
	head.next = B

	C := &Node{data: "C", next: nil, prev: B}
	B.next = C

	tail.data = "D"
	tail.next = head
	tail.prev = C
	C.next = tail

	head.prev = tail

	p := head

	add("I", 2)
	delete(2)
	for {
		fmt.Println(p.data)
		p = p.next

		if p == head {
			break
		}
	}
}

func add(data string, index int) {
	Elem := &Node{data: data, next: nil, prev: nil}

	p := head
	i := 0
	for {
		if p.next == nil || i+1 == index {
			break
		} 
		p = p.next
		i++
	}

	Elem.next = p.next
	p.next = Elem

	Elem.prev = p
	Elem.next.prev = Elem
}

func delete(index int) {
	p := head
	i := 0
	for {
		if p.next == nil || i+1 == index {
			break
		} 
		p = p.next
		i++
	}

	temp := p.next

	p.next = p.next.next
	p.next.prev = p

	temp.next = nil
	temp.prev = nil
}