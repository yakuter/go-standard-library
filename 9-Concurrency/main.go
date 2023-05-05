package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		defer wg.Done()
		for i := 0; i < 10; i++ {
			fmt.Printf("%d ", i)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	go func() {
		defer wg.Done()
		for c := 'a'; c <= 'z'; c++ {
			fmt.Printf("%c ", c)
			time.Sleep(100 * time.Millisecond)
		}
	}()

	wg.Wait()
}
