package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Node defines the basic building of the binary tree
type Node struct {
	// Data data for the node
	Data interface{}
	// Defines the sleep time during processing to simulate delay
	Sleep time.Duration
	// Left Child
	Left *Node
	// Right Child
	Right *Node
}

// NewNode returns a new *Node
func NewNode(data interface{}) *Node {

	node := new(Node)

	node.Data = data
	node.Left = nil
	node.Right = nil

	rand.Seed(time.Now().UTC().UnixNano())
	duration := int64(rand.Intn(100))

	// Random delay time set from 0-100 microseconds
	node.Sleep = time.Duration(duration) * time.Microsecond

	return node

}

// ProcessNode node processing function for normal DFS
func (n *Node) ProcessNode(searchNumber int) (int) {
	var hello []int
	
	for i := 0; i < 5000; i++ {
		time.Sleep(n.Sleep)
		hello = append(hello, i)
	}
	
	if _, ok:= found.found[searchNumber]; ok{
		fmt.Println("skip", n.Data)
		return 0
	}
	
	if n.Data.(int) == searchNumber{
		found.found[searchNumber] = searchNumber
		fmt.Println("found", searchNumber)
		return searchNumber
	}

	fmt.Printf("Node %v ✅\n", n.Data)
	return n.Data.(int)
}

// ProcessNodeParallel node processing function for parallel DFS
func (n *Node) ProcessNodeParallel(searchNumber int) (int) {

	defer wg.Done()

	var hello []int
	for i := 0; i < 5000; i++ {
		time.Sleep(n.Sleep)
		hello = append(hello, i)
	}

	if _, ok:= found.found[searchNumber]; ok{
		fmt.Println("skip", n.Data)
		return 0
	}

	if n.Data.(int) == searchNumber{
		found.mu.Lock()
		defer found.mu.Unlock()
		found.found[searchNumber] = searchNumber
		fmt.Println("found", searchNumber)
		return searchNumber
	}

	fmt.Printf("Node %v ✅\n", n.Data)
	return n.Data.(int)
}

// DFS normal DFS function
func (n *Node) DFS(searchNumber int) {

	if n == nil {
		return
	}

	n.Left.DFS(searchNumber)
	ans := n.ProcessNode(searchNumber)
	if ans == 0{
		return
	}
	n.Right.DFS(searchNumber)

}

// DFSParallel parallel DFS Function
func (n *Node) DFSParallel(searchNumber int) {

	// Removes goroutine from WaitGroup when function is returning
	defer wg.Done()

	if n == nil {
		return
	}

	// Adds a goroutine for left child DFS
	wg.Add(1)
	go n.Left.DFSParallel(searchNumber)

	// Adds a goroutine for Processing
	wg.Add(1)
	go n.ProcessNodeParallel(searchNumber)

	// Adds a goroutine for right child DFS
	wg.Add(1)
	go n.Right.DFSParallel(searchNumber)

}