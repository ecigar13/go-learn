package main

import (
	"gorgonia.org/tensor"
)

func createTensorArray(x int, y int) *tensor.Dense {
	ans := tensor.New(tensor.WithShape(x, y), tensor.WithBacking(tensor.Range(tensor.Float32, 0.0, x*y)))
	return ans
}
