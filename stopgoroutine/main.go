package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// Global WaitGroup initialised to keep track of goroutines in process
var wg sync.WaitGroup

type SafeCounter struct {
	mu sync.Mutex
	found  map[int]int
}
var found = SafeCounter{found: make(map[int]int)}

func main() {

	// Go will use maximum number of processors available to process goroutines
	processors := runtime.GOMAXPROCS(runtime.NumCPU())

	// Creates a basic Binary Tree
	//				1
	//		2				3
	//	4		5 		6		7

	root := NewNode(1)
	root.Left = NewNode(2)
	root.Right = NewNode(3)
	root.Left.Left = NewNode(4)
	root.Left.Right = NewNode(5)
	root.Right.Left = NewNode(6)
	root.Right.Right = NewNode(7)
	searchNumber:= 3

	// ============== WITHOUT GOROUTINES

	// Starts the timer
	start := time.Now()

	// Start depth first search on the root
	root.DFS(searchNumber)

	// DFS Finished
	fmt.Printf("\nTime elapsed: %v\n\n", time.Since(start))

	// =============== WITH GOROUTINES

	found = SafeCounter{found: make(map[int]int)}
	// Starts the timer
	start = time.Now()

	// Adds one goroutine the WaitGroup
	wg.Add(1)
	// Start the DFS Goroutine
	go root.DFSParallel(searchNumber)

	// Waits for all goroutines to complete
	wg.Wait()

	fmt.Printf("\nProcessors: %v Time elapsed: %v\n", processors, time.Since(start))

}