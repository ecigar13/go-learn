package main

import (
	"io/ioutil"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}
func readFile(path string) []byte {
	dat, err := ioutil.ReadFile(path)
	check(err)
	return dat
}
