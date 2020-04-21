package main

import (
	"fmt"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 10; i++ {
		wg.Add(1)
		i := i // fix the race condition
		go func() {
			fmt.Println(i)
			wg.Done()
		}()
	}
	wg.Wait()
}
