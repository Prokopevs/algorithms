package main
import (
	"fmt"
	"sync"
	
	)


func main() {
	var c1, c2, c3 = make(chan int, 3), make(chan int, 3), make(chan int, 3)

	for i:=0; i<3; i++ {
		c1 <- i
		c2 <- i
		c3 <- i
	}

	ch := crea(c1, c2, c3)

	count := len(ch)
	for i:=0; i<count; i++ {
		fmt.Println(<-ch)
	}
}

func crea(in ...<-chan int) chan int{
	var mu sync.Mutex
	var wg sync.WaitGroup 
	wg.Add(3) 
	ch := make(chan int, 9)

	for i:= 0; i<len(in); i++ {
		
		go func(h <-chan int) {
			count := len(h)
			for i:=0; i<count; i++ {
				mu.Lock()
				ch <- <-h
				mu.Unlock()
			}
			
			wg.Done()
		}(in[i])
	}

	wg.Wait() 
	return ch
}







