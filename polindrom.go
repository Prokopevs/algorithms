// package main

// import (
// 	"fmt"
// 	"strings"
// 	"reflect"
// )

// func main() {
// 	str := "анна"

// 	r := strings.Split(str, "")
// 	b := true

// 	for i := 0; i < len(r)/2; i++ {
// 		if r[i] != r[len(r)-1-i] {
// 			b = false
// 		}
// 	}

// 	fmt.Println(b)
// }

// func main() {
// 	str := "anna"
// 	r := str
// 	b := true

// 	for i := 0; i < len(r)/2; i++ {
// 		if r[i] != r[len(r)-1-i] {
// 			b = false
// 		}
// 	}

// 	fmt.Println(b)
// }

// func main() {
// 	str := "привет"

// 	reversedStr := strings.Builder{}
// 	b := true

// 	for i := len(str)-1; i >= 0; i-- {
// 		reversedStr.WriteByte(str[i])
// 	}

// 	if str != reversedStr.String() {
// 		b = false
// 	}
// 	fmt.Println(b)
// }




