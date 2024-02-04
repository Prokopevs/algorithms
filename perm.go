// package main

// import (
// 	"fmt"
// )

// // Perm вызывает f с каждой пермутацией a.
// func Perm(a []rune) {
//   perm(a, 0)
// }

// // Пермутируем значения в индексе i на len(a)-1.
// func perm(a []rune, i int) {
//   if i > len(a) {
//     // for _, v := range a {
//     //     fmt.Print(string(v))
//     // }
//     // fmt.Println()
//     return
//   }
//   perm(a, i+1)
//   for _, v := range a {
// 	fmt.Print(string(v))
//   }
//   fmt.Println(i+1)

//   for j := i + 1; j < len(a); j++ {
//     a[i], a[j] = a[j], a[i]
// 	for _, v := range a {
//         fmt.Print(string(v))
//     }
// 	fmt.Println("-----------", i+1)
//     perm(a, i+1)
//     a[i], a[j] = a[j], a[i]
//   }
// }

// func main() {
//   Perm([]rune("abc"))
// }