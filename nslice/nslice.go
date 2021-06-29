package main

import (
	"fmt"
)

func create2DSlice(x int, y int) [][]int {
	ans := [][]int{}
	for i := 0; i < x; i++ {
		ans = append(ans, []int{})
		for j := 0; j < y; j++ {
			ans[i] = append(ans[i], j)
		}
	}
	return ans
}

func main() {
	a := createTensorArray(2, 2)
	fmt.Println(a)
	b := create2DSlice(3, 3)
	fmt.Println(b, len(b), cap(b))
}
