package main

import (
	"fmt"
	"strings"
)


func main() {
	str := "anna"
	reversedArr := make([]string, len(str))
	
	for i:=0; i<len(str); i++ {
		reversedArr[i] = string(str[len(str)-i-1])
	}

	fmt.Println(str)
	fmt.Println(strings.Join(reversedArr, ""))
	
}




