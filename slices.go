package main 

import (
	"fmt"
)

func main () {

	// удалить с конца
	// slice := []int{1, 3, 4, 9, 4}
	// slice = slice[0:len(slice)-1]
	// fmt.Println(slice)


	// удалить с начала
	// slice2 := []int{1, 3, 4, 9, 4}
	// slice2 = slice2[1:] 
	// fmt.Println(slice2)

	// удалить с середины
	// slice3 := []int{1, 3, 4, 9, 4}
	// n := 3
	// slice3C := append(slice3[0:n], slice3[n+1:]...)
	// fmt.Println(slice3C)

	// вставка в начало
	// slice4 := []int{1, 3, 4, 9, 4}
	// n := 0
	// slice4C := append(slice4[:n+1], slice4[n:]...)
	// slice4C[0] = 10
	// fmt.Println(slice4C)


	nums := make([]int, 1, 2)
	fmt.Println(nums)

	// appendSlice(nums, 1024)
	// fmt.Println(nums[:2])

	// mutateSlice(nums, 0, 512)
	// fmt.Println(nums)

	a := append(nums, 1)
	fmt.Println(a)
	fmt.Println(nums[:2])


}

func appendSlice (sl []int, val int) {
	sl = append(sl, val)
}

func mutateSlice (sl []int, idx, val int) {
	sl[idx] = val
}




