package main

import (
	"sync"
	"fmt"
)

var COUNTER int 

func main() {
	var wg sync.WaitGroup
	var lock sync.Mutex
	for i := 0; i < 1000; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			lock.Lock()
			COUNTER ++
			lock.Unlock()
		}()
	}
	wg.Wait()
	fmt.Println(COUNTER)
}