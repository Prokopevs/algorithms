// package main

// import (
// 	"fmt"
// 	"strings"
// )

// func main() {
//     word := "helloaa"
// 	wordToLower := strings.ToLower(word)
// 	count := 0

// 	for i := 0; i < len(wordToLower); i++ {
// 		switch string(wordToLower[i]) {
// 		case "a", "e", "i", "o", "u", "y" :
// 			count++
// 		}

// 	}

// 	// for _, v := range wordToLower {
// 	// 	switch v {
// 	// 	case 'a', 'e', 'i', 'o', 'u', 'y' :
// 	// 		count++
// 	// 	}
// 	// }

// 	vow := []string{"a", "e", "i", "o", "u", "y"}
// 	c := 0

// 	for _, v := range wordToLower {
// 		if strings.ContainsAny(string(v), strings.Join(vow, "")) {
// 			c++
// 		}
// 	}

// 	fmt.Println(count)
// 	fmt.Println(c)
// }