// package main

// import (
// 	"fmt"
// )

// type Node struct {
// 	data string
// 	next *Node
// }

// var head *Node = &Node{data: "LU", next: nil}
// var tail *Node = new(Node)

// func main() {
// 	var ff *Node = &Node{data: "GW", next: nil}
// 	head.next = ff

// 	var gg *Node = &Node{data: "RT", next: nil}
// 	ff.next = gg

// 	tail = &Node{data: "HH", next: nil}
// 	gg.next = tail

// 	add("JJ")

// 	p := head
// 	for {
// 		fmt.Println(p.data)
// 		if p.next == nil {
// 			break
// 		}
// 		p = p.next
// 	}
// }

// func add (data string) {
// 	var newN *Node = &Node{data: data, next: nil}
// 	tail.next = newN

// 	tail = newN
// }


