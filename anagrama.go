// package main

// import (
// 	"fmt"
// 	// "sort"
// 	// "strings"
// 	// "unicode/utf8"
// )

// func main() { 
// 	str1 := "anna"
// 	str2 := "aann"

// 	if len(str1) != len(str2) {
// 		fmt.Println("false")
// 		return 
// 	}

// 	// // if utf8.RuneCountInString(str1) != utf8.RuneCountInString(str2) {
// 	// // 	fmt.Println("false")
// 	// // 	return 
// 	// // }

// 	// str1Arr := strings.Split(str1, "") 
// 	// str2Arr := strings.Split(str2, "") 
// 	// sort.Strings(str1Arr)
// 	// sort.Strings(str2Arr)

// 	// for i := 0; i < len(str1Arr)/2; i++ {
// 	// 	if str1Arr[i] != str2Arr[i] {
// 	// 		fmt.Println("false here")
// 	// 		return 
// 	// 	}

// 	// 	if str1Arr[len(str1Arr)-1-i] != str2Arr[len(str1Arr)-1-i] {
// 	// 		fmt.Println("fal")
// 	// 		return 
// 	// 	}
// 	// }

// 	// fmt.Println("true")
// 	// return 

// 	// -------------------------------- //
// 	hashTable := map[byte]int{}
// 	// hashTable2 := make(map[string]int)

// 	for i := 0; i < len(str1); i++ {
// 		hashTable[str1[i]]++
// 		hashTable[str2[i]]++
// 	}

// 	for _, v := range hashTable {
// 		if v % 2 != 0 {
// 			fmt.Println("false")
// 			return
// 		}
// 	}

// 	fmt.Println("true")
// 	return 
// } 