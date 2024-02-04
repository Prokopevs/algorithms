// package main

// import (
// 	"fmt"
// )

// func main () {
// 	slice := []int{1, 3, 4, 9, 4, 7, 8}

// 	numIndex := len(slice)/2
// 	fmt.Println(numIndex)

// 	for i:= 0; i<numIndex; i++ {
// 		slice[i], slice[len(slice)-i-1] = slice[len(slice)-i-1], slice[i]

// 	}
// 	fmt.Println(slice)
// }

// func main() {
// 	slice := []int{1, 3, 4, 9, 4, 7, 8}
// 	reverceSlice := make([]int, 0, len(slice))

// 	// for _, v := range slice {
// 	// 	reverceSlice = append(reverceSlice[:1], reverceSlice[:]...)
// 	// 	reverceSlice[0] = v
// 	// }

// 	for i := len(slice)-1; i >= 0; i-- {
// 		reverceSlice = append(reverceSlice, slice[i])
// 	}

// 	fmt.Println(reverceSlice)
// }
