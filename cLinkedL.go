// package main 

// import (
// 	"fmt"
// )


// type Node struct {
// 	data string
// 	next *Node
// }

// 	var head *Node = new(Node)
// 	var tail *Node = new(Node)

// func main() {

// 	head.data = "A"
// 	head.next = nil

// 	var B *Node = &Node{data: "B", next: nil}
// 	head.next = B  

// 	var C *Node = &Node{data: "C", next: nil}
// 	B.next = C 

// 	tail.data = "D"
// 	tail.next = head
// 	C.next = tail

// 	p := head

// 	add("E", 2)
// 	delete(2)

// 	for {
// 		fmt.Println(p.data)
// 		p = p.next

// 		if p == head {
// 			fmt.Println(p.data)
// 			break
// 		}
// 	}
// }

// func add(data string, index int) {
// 	p := head
// 	i := 0

// 	E := &Node{data, nil}

// 	for {
// 		if p.next == nil || i+1 >= index {
// 			break
// 		}
// 		p = p.next
// 		i++
// 	}

// 	E.next = p.next
// 	p.next = E
// }

// func delete(index int) {
// 	p := head
// 	i := 0

// 	for {
// 		if p.next == nil || i+1 >= index {
// 			break
// 		}
// 		p = p.next
// 		i++
// 	}

// 	temp := p.next
// 	p.next = p.next.next

// 	temp.next = nil
// }

