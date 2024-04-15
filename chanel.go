// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch1 := make(chan int)
// 	ch2 := make(chan int, 5)


// 	ch3 := make(chan int)
// 	go func() {
// 		for i:=0; i<5; i++ {
// 			ch1 <- i
// 		}
// 		close(ch1)
// 	}()

// 	go func() {
// 		for i := range ch1 {
// 			ch2 <- i*i
// 		}
// 		close(ch3)
// 	}()

// 	<-ch3

// 	for i:=0; i<5; i++ {
// 		fmt.Println(<-ch2)
// 	}
// }
