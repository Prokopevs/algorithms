// package main

// import "fmt"

// func main() {
// 	slice := []int{1, 4, 3, 7, 2, 1, 9, 2}

// 	for i := 0; i < len(slice); i++ {
// 		for j := 0; j < len(slice)-1-i; j++ {
// 			if slice[j] > slice[j+1] {
// 				slice[j], slice[j+1] = slice[j+1], slice[j]
// 			}
// 		}
// 	}

// 	// for i := 0; i < len(slice); i++ {
// 	// 	for j := 0; j < len(slice)-1-i; j++ {
// 	// 		if slice[j] > slice[j+1] {
// 	// 			temp := slice[j]
// 	// 			slice[j] = slice[j+1]
// 	// 			slice[j+1] = temp
// 	// 		}
// 	// 	}
// 	// }

// 	fmt.Println(slice)
// }
