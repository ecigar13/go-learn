package main

import "fmt"

// fibonacci is a function that returns
// a function that returns an int.
func fibonacci() func() int {
	a := -1
	b := 0
	return func() int {
		if a == -1 && b == 0 {
			a = 0
			return a
		}
		if a == 0 && b == 0 {
			b = 1
			return b
		}
		temp := a + b
		a = b
		b = temp
		return temp
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
