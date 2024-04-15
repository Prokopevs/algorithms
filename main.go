// package main 

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{1, 3, 4, 6, 8, 10, 55, 56, 59, 70, 79, 81, 91, 10001}

// 	min := 0
// 	max := len(arr)-1
// 	find := 81

// 	for i:=0; i<len(arr); i++ {
// 		center := (min+max)/2

// 		if arr[center] == find {
// 			fmt.Println(true, center)
// 			return 
// 		}
// 		if arr[center] > find {
// 			max = center - 1
// 			continue
// 		} 
// 		if arr[center] < find {
// 			min = center + 1
// 		}
// 	}
	
// }