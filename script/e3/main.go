package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	ch := make(chan int)

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for {
				if v, ok := <-ch; ok {
					fmt.Println(v)
				} else {
					break
				}
			}
		}()
	}

	for i := 0; i < 10; i++ {
		ch <- i
	}

	close(ch)
	wg.Wait()
}
