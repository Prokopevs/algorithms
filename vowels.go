package main

import (
	"fmt"
	"strings"
)

func main() {
    word := "helloaa"
	wordToLower := strings.ToLower(word)
	count := 0

	for i := 0; i < len(wordToLower); i++ {
		switch string(wordToLower[i]) {
		case "a", "e", "i", "o", "u", "y" :
			count++
		}

	}

	// for _, v := range wordToLower {
	// 	switch v {
	// 	case 'a', 'e', 'i', 'o', 'u', 'y' :
	// 		count++
	// 	}
	// }

	




	fmt.Println(count)
}