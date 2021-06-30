package main

import (
	"fmt"
	"strconv"
	"sync"
)

// sync.WaitGroup eliminates the use of time.sleep
// the reason is: goroutines returns immediately without execution.
// time.sleep prevents it, but we don't know how long to wait
// WaitGroup.Wait() waits until all functions return. We can handle to error and return it.
func say(s string, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := 0; i < 2; i++ {
		fmt.Println(s)
	}
}

func main() {
	var wg sync.WaitGroup
	for i := 1; i <= 3; i++ {
		wg.Add(1)
		go say(strconv.Itoa(i), &wg)
	}
	wg.Wait()
}
