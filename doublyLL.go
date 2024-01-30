// package main

// import "fmt"

// type Node struct {
// 	data string
// 	prev *Node
// 	next *Node
// }

// 	var head *Node = new(Node)
// 	var tail *Node = new(Node)
// func main() {
	
// 	head.data = "AA"
// 	head.next = nil
// 	head.prev = nil

// 	var holand *Node = &Node{data: "BB", prev: head, next: nil}
// 	head.next = holand

// 	var ronaldo *Node = &Node{data: "CC", prev: holand, next: nil}
// 	holand.next = ronaldo

// 	tail.data = "GG"
// 	tail.prev = ronaldo
// 	tail.next = nil
// 	ronaldo.next = tail

// 	// add to the end
// 	add("MM")
// 	insert("ST", 2)
// 	delete(2)

// 	p := head
// 	for {
// 		fmt.Println(p.data)
// 		if p.next == nil {
// 			break
// 		} 
// 		p = p.next
// 	}
// }

// func add (data string ) {
// 	var messi *Node = &Node{data: data, prev: tail, next: nil}
// 	tail.next = messi
// 	tail = messi
// }

// func insert (data string, index int) {
// 	var p = head
// 	var i = 0

// 	for {
// 		if p.next == nil || i >= index-1 {
// 			break
// 		}
// 		p = p.next
// 		i++
// 	}

// 	var newEl *Node = &Node{data: data, prev: nil, next: nil}

// 	newEl.next = p.next 
// 	p.next = newEl

// 	newEl.prev = p
// 	newEl.next.prev = newEl
// }

// func delete ( index int) {
// 	var p = head
// 	var i = 0

// 	for {
// 		if p.next == nil || i >= index-1 {
// 			break
// 		}
// 		p = p.next
// 		i++
// 	}

// 	temp := p.next
// 	p.next = p.next.next
// 	p.next.prev = p

// 	temp.next = nil
// 	temp.prev = nil
// }