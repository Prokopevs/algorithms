// package main

// import "fmt"

// type stack struct {
// 	items []int
// }

// func (s *stack) push(item int) {
// 	s.items = append(s.items, item)
// }

// func (s *stack) pop() int {
// 	var lastItem int
// 	if len(s.items) != 0 {
// 		lastItem = s.items[len(s.items)-1]
// 	} else {
// 		return -1
// 	}
	
// 	s.items = s.items[:len(s.items)-1]

// 	return lastItem
// }

// func newStack() *stack {
// 	return &stack{items: []int{}}
// }

// func main() {
// 	stack := newStack()

// 	stack.push(1)
// 	stack.push(2)
// 	elem := stack.pop()
// 	stack.pop()
// 	stack.pop()

// 	fmt.Println(stack)
// 	fmt.Println(elem)
// }