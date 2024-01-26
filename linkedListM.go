// package main

// import (
// 	"fmt"
// )

// type Node struct {
// 	data string
// 	next *Node
// }


// var tail *Node = new(Node)


// func main() {
// 	var head *Node = &Node{data: "LU", next: nil}
// 	tail = head

// 	tail = tail.add("GW", tail)
// 	tail = tail.add("OM", tail)

// 	head.insert("ie", 1)

// 	head.delete(2)

// 	p := head
// 	for {
// 		fmt.Println(p.data)
// 		if p.next == nil {
// 			break
// 		}
// 		p = p.next
// 	}
// }


// func (s *Node) add(data string, tail *Node) (t *Node) {
// 	var newN *Node = &Node{data: data, next: nil}
// 	tail.next = newN
// 	tail = newN

// 	return tail
// } 

// func (s *Node) insert(data string, index int) {
// 	var p = s
// 	var i = 0

// 	for {
// 		if p.next == nil || i >= index-1 {
// 			break
// 		}
// 		p = p.next
// 		i++
// 	}


// 	var newN *Node = &Node{data: data, next: p.next}
// 	p.next = newN
// } 

// func (s *Node) delete(index int) {
// 	var p = s
// 	var i = 0

// 	for {
// 		if p.next == nil || i >= index-1 {
// 			break
// 		}
// 		p = p.next
// 		i++
// 	}

// 	tempD := p.next

// 	p.next = p.next.next

// 	tempD.next = nil
// } 


