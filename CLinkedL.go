package main 

import (
	"fmt"
)


type Node struct {
	data string
	next *Node
}

func main() {

	var head *Node = new(Node)
	var tail *Node = new(Node)

	head.data = "A"
	head.next = nil

	var B *Node = &Node{data: "B", next: nil}
	head.next = B  

	var C *Node = &Node{data: "C", next: nil}
	B.next = C 

	tail.data = "D"
	tail.next = head
	C.next = tail

	p := head
	for {
		fmt.Println(p.data)
		p = p.next

		if p == head {
			fmt.Println(p.data)
			break
		}
	}
}

