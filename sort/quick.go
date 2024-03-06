// package main 

// import (
// 	"fmt"
// )

// func main() {
// 	arr := []int{3, 4, 2, 1, 7, 8, 3, 10}

// 	fmt.Println(quick(arr))
// }

// func quick(arr []int) []int {
// 	if len(arr) < 2 {
// 		return arr
// 	}
// 	pivot := arr[0]
// 	var greater, less []int

// 	for _, v := range arr[1:] {
// 		if v > pivot {
// 			greater = append(greater, v)
// 		} else {
// 			less = append(less, v)
// 		}
// 	}

// 	result := append(quick(less), pivot)
// 	result = append(result, quick(greater)...)

// 	return result
// }