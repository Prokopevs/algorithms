package main 

import (
	"fmt"
)

func main () {
	array := [5]int{1, 3, 4, 9, 4}

	for i := 0; i < len(array); i++ {
		minIndex := i
		for j := i; j < len(array)-1; j++ {
			if array[minIndex] > array[j+1] {
				minIndex = j+1
			}
		}

		if i != minIndex {
			temp := array[i]
			array[i] = array[minIndex]
			array[minIndex] = temp
		}
		
	}

	fmt.Println(array)
}