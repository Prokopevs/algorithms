package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
	"sync"
)

func main() {
	b, err := ioutil.ReadFile("text.txt")
    if err != nil {
        panic(err)
    }
	str := string(b)

	arr := strings.Split(str, "\n")

	wg := sync.WaitGroup{}

	for _, v := range arr {
		wg.Add(1)
		go func(url string) {
			resp, err := http.Get(url)
			if err != nil {
				fmt.Println("error")
			}
			fmt.Println(resp.Status[3:], url)
			wg.Done()
		}(v)
	}

	wg.Wait()
}