// package main 

// import (
// 	"fmt"
// )

// func main() {
// 	array := [5]int{1, 3, 4, 9, 4}

// 	for i := 0; i < len(array); i++ {
// 		minIndex := i
// 		for j := i; j < len(array)-1; j++ {
// 			if array[minIndex] > array[j+1] {
// 				minIndex = j+1
// 			}
// 		}

// 		if i != minIndex {
// 			temp := array[i]
// 			array[i] = array[minIndex]
// 			array[minIndex] = temp
// 		}
		
// 	}

// 	fmt.Println(array)
// }


// func main() {
// 	slice := []int{1, 4, 3, 7, 2, 1, 9, 2}

// 	for i := 0; i < len(slice); i++ {
// 		minIndex := i
// 		for j := i; j < len(slice)-1; j++ {
// 			if slice[minIndex] > slice[j+1] {
// 				minIndex = j+1
// 			}
// 		}

// 		if minIndex != i {
// 			slice[i], slice[minIndex] = slice[minIndex], slice[i]
// 		}	
// 	}
// 	fmt.Println(slice)
// }

