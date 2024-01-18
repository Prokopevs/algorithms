// package main 

// import (
// 	"fmt"
// )

// func main () {
// 	array := [5]int{1, 3, 4, 9, 4}

// 	for i := 0; i < len(array); i++ {
// 		for j := 0; j < len(array)-1-i; j++ {
// 			if array[j] > array[j+1] {
// 				temp := array[j]
// 				array[j] = array[j+1]
// 				array[j+1] = temp
// 			}
// 		}
// 	}

// 	fmt.Println(array)
// }