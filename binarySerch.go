package main 

import (
	"fmt"
	
)

func main () {
	slice := []int{1, 3, 4, 5, 6, 7, 8}
	find := 5
	min := 0
	high := len(slice)-1

	res := false
	for i := 0; i < len(slice); i++ { 
		mid := (min + high)/2
		if find == slice[mid] {
			res = true
			break
		}
		if find > slice[mid] {
			min = mid + 1  
			continue 
		}
		if find < slice[mid] {
			high = mid -1
			continue
		} 
	}
	fmt.Println(res)
}

