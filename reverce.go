package main 

import (
	"fmt"
	
)

func main () {
	slice := []int{1, 3, 4, 9, 4,7}

	var value int
	fmt.Scan(&value)
	var num int
	if len(slice)%2 == 0 {
		num = len(slice)/2
	} else {
		num = (len(slice)-1)/2
	}

	for i:= 0; i<num; i++ {
		temp := slice[i]
		slice[i] = slice[len(slice)-i-1]
		slice[len(slice)-i-1] = temp
	}
	fmt.Println(slice)
	fmt.Println(value)

}

