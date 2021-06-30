package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

type Tree struct {
	Left  *Tree
	Value int
	Right *Tree
}

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	defer close(ch)
	var walker func(t *tree.Tree)
	walker = func(t *tree.Tree) {
		if t == nil {
			return
		}
		walker(t.Left)
		ch <- t.Value
		walker(t.Right)
	}
	walker(t)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch := make(chan int)
	go Walk(t1, ch)
	ch1 := make(chan int)
	go Walk(t2, ch1)

	for {
		v1, ok1 := <-ch
		v2, ok2 := <-ch1
		if v1 != v2 || ok1 != ok2 {
			return false
		}
		if !ok1 {
			break
		}
	}
	return true
}

func main() {
	fmt.Println("1 and 1 same", Same(tree.New(1), tree.New(1)))
	fmt.Println("1 and 1 same", Same(tree.New(1), tree.New(2)))
	fmt.Println(tree.New(1))
}
